package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Get user's home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Assign the location for the go-chat logs
	// TODO: Allow an XML/YAML/JSON file to be parsed for this data
	var logpath string = usr.HomeDir + "/.go-chat/log"
	// Check if logpath exists
	_, err = os.Stat(logpath)
	if err != nil {
		// TODO: Add this to "syslog" file
		fmt.Println("go-chat log path doesn't exist... Creating it at: ", logpath)
		os.MkdirAll(logpath, 660)
	}
	// Used by other parts of go-chat
	os.Setenv("LOGPATH", logpath)

	router := httprouter.New()
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)
	router.GET("/c/:channel", loadChannel)
	router.GET("/", loadRoot)

	log.Fatal(http.ListenAndServe(":8080", router))
}

/****** Me implementing the ServeHTTP method such as to use
http.Handle(PATH string, SOMEFUNC function)
type fooHandler func(http.ResponseWriter, *http.Request)

var fooFunction fooHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, foo")
}

func (f fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
*****/
