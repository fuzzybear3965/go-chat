package main

import (
	"fmt"
	_ "github.com/fuzzybear3965/go-chat/static"
	//	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type channelInfo struct {
	ChannelName string
	Body        []byte
	LogPath     string
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
			fmt.Println("Could not create file ", logPath, " when trying to create it because it doesn't exist.")
			log.Fatal(err)
		} else {
			fh.Write([]byte("History"))
			if _, err := fh.Stat(); err != nil {
				fmt.Println("Could not get file info.")
				log.Fatal(err)
			} else {
				if err := os.Chmod(logPath, 660); err != nil {
					fmt.Println("Could not change permissions for file ", logPath)
				}
				if err := fh.Close(); err != nil {
					fmt.Println("Could not close the file for channel ", channelName, ": ", logPath)
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

	fmt.Println("Channel", ci.ChannelName, "has been requested.")

	// Read template for the chat page.
	t := channelTemplate
	//if err != nil {
	//	// TODO: Add the below to a log file.
	//	fmt.Println("Could not read the channel template.")
	//	log.Fatal(err)
	//}

	t.Execute(w, ci)
}

func saveChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Save the state of the channel to a file
	// TODO: Do this only if there are no other clients on the channel.
	// Otherwise we'll keep the unsaved data in a <chan string>.
	ci := getChannelInfo(r.URL.Path, params)

	fmt.Println("The state of channel", ci.ChannelName, "is going to be saved.")

	// Get set data, if any
	r.ParseForm()
	// Get the message sent
	msgStringSlice := r.Form["message"]
	// Prepare to cenvert []string to -> []byte
	var msgBytes []byte = ci.Body
	// Add a new line to separate this message from other messages.
	msgBytes = append(msgBytes, '\n')
	for _, val := range msgStringSlice {
		msgBytes = append(msgBytes, []byte(val)...)
	}

	ci.Body = msgBytes

	err := ci.channelLogSave()

	if err != nil {
		fmt.Println("Problem saving channel ", ci.ChannelName)
		log.Fatal(err)
	} else {
		fmt.Println("Saved the state of channel", ci.ChannelName, "just fine.")
	}
	// Reload the channel for the user.
	loadChannel(w, r, params)
}
