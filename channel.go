package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type channelInfo struct {
	ChannelName string
	LogPath     string
}

// Create a struct to hold the JS and CSS templates
type templateAssets struct {
	CSS string
	JS  template.JS
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
func (s *serverContext) getChannelInfo(urlpath string, routerparams httprouter.Params) *channelInfo {
	// Parse urlpath into different /foo/bar -> ["foo","bar"]
	pathArray := strings.Split(urlpath, "/")
	channelName := pathArray[len(pathArray)-1]
	if channelName != routerparams.ByName("channel") {
		fmt.Println("Problem parsing channel name.")
		log.Fatal(nil)
	}

	logName := channelName + ".txt"
	logPath := s.logdir + "/" + logName

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
