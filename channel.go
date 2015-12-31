package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Channel struct {
	Title string
	Body  []byte
}

func channelOpen(title string) (*Channel, error) {
	fileName := title + ".txt"
	filebytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Channel{Title: title, Body: filebytes}, err
}

func (p *Channel) channelSave() error {
	pathArray := strings.Split(p.Title, "/")
	channelName := pathArray[len(pathArray)-1]
	fileName := channelName + ".txt"

	fmt.Println("Saving the channel ", channelName, " to ", os.Getenv("LOGPATH")+"/"+fileName, ".")
	return ioutil.WriteFile(os.Getenv("LOGPATH")+"/"+fileName, p.Body, 0600)
}

func loadChannel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Welcome to the %v channel.", channelName)
	outputstring := fmt.Sprintf("Welcome to the %q page.", path)
	page := &Channel{Title: path, Body: []byte(outputstring)}
	err := page.channelSave()
	if err != nil {
		fmt.Println("Could not save the channel.")
		log.Fatal(err)
	}
}
