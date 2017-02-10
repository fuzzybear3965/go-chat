//go:generate go-bindata -o assets.go assets templates
//go:generate goversioninfo
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var cookieStore = sessions.NewCookieStore([]byte("secret passphrase"))
var src = rand.NewSource(time.Now().UnixNano())

func main() {
	// Get user's home directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Couldn't obtain user information.")
		log.Fatal(err)
	}

	// Assign the location for the go-chat logs
	// TODO: Allow an XML/YAML/JSON file to be parsed for this data
	var logpath = usr.HomeDir + "/.go-chat/log"
	// Check if logpath exists
	_, err = os.Stat(logpath)
	if err != nil {
		// TODO: Add this to "syslog" file
		fmt.Println("go-chat log path doesn't exist... Creating it at: ", logpath)
		err = os.MkdirAll(logpath, 0660)
		if err != nil {
			fmt.Println("Couldn't make go chat log at: ", logpath)
			log.Fatal(err)
		}
	}
	sc := &serverContext{port: 80, users: nil, logdir: logpath}

	router := httprouter.New()
	router.GET("/login", sc.getLogin)
	router.POST("/login", sc.postLogin)
	router.GET("/c/:channel", sc.loadChannel)
	router.POST("/c/:channel", sc.saveChannel)
	// Add route for root
	router.GET("/", sc.rootHandler)
	// Add js, css handler
	router.ServeFiles("/assets/*filepath", http.Dir("./assets"))

	log.Fatal(http.ListenAndServe(":80", router))
}
