package main

import (
	"fmt"
	"github.com/fuzzybear3965/go-chat/static"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type channelInfo struct {
	ChannelName string
	Body        []byte
	LogPath     string
}

// Create a struct to hold the JS and CSS templates
type templateScripts struct {
	ChannelTmplJS  template.JS
	ChannelTmplCSS template.CSS
}

// Open the channel log and return the a channelInfo object
func getChannelInfo(urlpath string, routerparams httprouter.Params) *channelInfo {
	// Parse urlpath into different /foo/bar -> ["foo","bar"]
	pathArray := strings.Split(urlpath, "/")
	channelName := pathArray[len(pathArray)-1]
	if channelName != routerparams.ByName("channel") {
		fmt.Println("Problem parsing channel name.")
		log.Fatal(nil)
	}

	logName := channelName + ".txt"
	logPath := os.Getenv("LOGPATH") + "/" + logName

	// Check if filePath exists
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		fmt.Println(channelName, "is a new channel. Making the following file:", logPath)
		if fh, err := os.Create(logPath); err != nil {
			fmt.Println("Could not create file", logPath, "(trying to create it because it doesn't exist).")
			log.Fatal(err)
		} else {
			if _, err := fh.Stat(); err != nil {
				fmt.Println("Could not get file info.")
				log.Fatal(err)
			} else {
				if err := os.Chmod(logPath, 660); err != nil {
					fmt.Println("Could not change permissions for file", logPath)
				}
				if err := fh.Close(); err != nil {
					fmt.Println("Could not close the file for channel", channelName, ":", logPath)
					log.Fatal(err)
				}
			}
		}
	}

	logbytes, err := ioutil.ReadFile(logPath)
	if err != nil {
		fmt.Println("Could not get channel info.")
		log.Fatal(err)
		return nil
	}
	return &channelInfo{ChannelName: channelName, Body: logbytes, LogPath: logPath}
}

func (p *channelInfo) channelLogSave() error {
	fmt.Println("Saving the channel", p.ChannelName, "log file to", p.LogPath, ".")
	return ioutil.WriteFile(p.LogPath, p.Body, 0600)
}

func loadChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get the complete URL path into an array
	ci := getChannelInfo(r.URL.Path, params)
	scriptTemplate := &templateScripts{
		ChannelTmplJS:  static.ChannelTmplJS,
		ChannelTmplCSS: static.ChannelTmplCSS,
	}

	fmt.Println("Channel", ci.ChannelName, "has been requested.")

	data := struct {
		Channel  *channelInfo
		Template *templateScripts
	}{
		ci,
		scriptTemplate,
	}

	static.ChannelTemplate.Execute(w, data)
}

func saveChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Save the state of the channel to a file
	// TODO: Do this only if there are no other clients on the channel.
	// Otherwise we'll keep the unsaved data in a <chan string>.
	ci := getChannelInfo(r.URL.Path, params)

	// Get set data, if any
	r.ParseForm()
	// Get the message sent
	msgStringSlice := r.Form["message"]

	if len(msgStringSlice[0]) == 0 {
		fmt.Println("Null message. Nothing being done.")
	} else {
		// Prepare to convert msgStringSlice []string to -> []byte
		var msgBytes []byte = ci.Body
		// Ensure Body is non-empty.
		if len(msgBytes) != 0 {
			// Add a new line to separate this message from other messages.
			msgBytes = append(msgBytes, '\n')
		} else {
			fmt.Println("This is the first line!")
		}

		for _, val := range msgStringSlice {
			msgBytes = append(msgBytes, []byte(val)...)
		}

		fmt.Println("The state of channel", ci.ChannelName, "is going to be saved.")

		ci.Body = msgBytes

		err := ci.channelLogSave()

		if err != nil {
			fmt.Println("Problem saving channel ", ci.ChannelName)
			log.Fatal(err)
		} else {
			fmt.Println("Saved the state of channel", ci.ChannelName, "just fine.")
		}
	}
	// Reload the channel for the user.
	loadChannel(w, r, params)
}
