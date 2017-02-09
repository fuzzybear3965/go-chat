package main

import (
	"fmt"
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
	LogPath     string
}

// Create a struct to hold the JS and CSS templates
type templateScripts struct {
	ChannelTmplJS  template.JS
	ChannelTmplCSS template.CSS
}

type messageContainer struct {
	message string
	userID  string
}

// return the history of the channel at any time for the user
func getLog(ci *channelInfo) string {
	bytes, err := ioutil.ReadFile(ci.LogPath)
	if err != nil {
		fmt.Println("Could not read", ci.LogPath)
	}
	return string(bytes)
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
				if err := os.Chmod(logPath, 0660); err != nil {
					fmt.Println("Could not change permissions for file", logPath)
				}
				if err := fh.Close(); err != nil {
					fmt.Println("Could not close the file for channel", channelName, ":", logPath)
					log.Fatal(err)
				}
			}
		}
	}

	return &channelInfo{ChannelName: channelName, LogPath: logPath}
}

func saveMessage(mc *messageContainer, ci *channelInfo) error {
	fmt.Println("Saving the message to the log file", ci.LogPath, ".")
	fh, err := os.OpenFile(ci.LogPath, os.O_APPEND|os.O_WRONLY, 0660)
	defer fh.Close()
	if err != nil {
		fmt.Println("Could not open logfile/history located at:", ci.LogPath)
		log.Fatal(err)
	}
	_, err = fh.WriteString(mc.userID + ": " + mc.message + string('\n'))
	return err
}

func loadChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get the complete URL path into an array
	ci := getChannelInfo(r.URL.Path, params)
	scriptTemplate := &templateScripts{
		ChannelTmplJS:  static.ChannelTmplJS,
		ChannelTmplCSS: static.ChannelTmplCSS,
	}

	fmt.Println("Channel", ci.ChannelName, "has been requested.")

	log := getLog(ci)
	data := struct {
		ChannelName string
		ChannelLog  string
		Template    *templateScripts
	}{
		ChannelName: ci.ChannelName,
		ChannelLog:  log,
		Template:    scriptTemplate,
	}

	static.ChannelTemplate.Execute(w, data)
}

func saveChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Save the state of the channel to a file
	// TODO: This needs to work with WebSockets so that new messages are pushed
	// to the client. Right now the user has to refresh the page.
	ci := getChannelInfo(r.URL.Path, params)
	s := getSession(w, r)

	// Get set data, if any
	r.ParseForm()

	var userID string = s.Values["userID"].(string)

	if len(r.Form["message"][0]) == 0 {
		fmt.Println("Null message. Nothing being done.")
	} else {
		// TODO: Add timestamp information below
		var msgString = &messageContainer{
			message: strings.Join(r.Form["message"], " "),
			userID:  userID,
		}

		fmt.Println("New message by user: ", msgString.userID, ". The state of channel", ci.ChannelName, "is going to be saved to", ci.LogPath, ".")

		err := saveMessage(msgString, ci)

		if err != nil {
			fmt.Println("Problem saving channel", ci.ChannelName)
			log.Fatal(err)
		} else {
			fmt.Println("Saved the state of channel", ci.ChannelName, "just fine.")
		}
	}
	// Reload the channel for the user.
	loadChannel(w, r, params)
}
