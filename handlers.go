package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// /
func (s *serverContext) rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	accessLog(r)
	data, err := Asset("templates/root.html")
	if err != nil {
		fmt.Println("Error acquiring root.html asset.")
	}
	root_template := template.New("root")
	root_template.Parse(string(data))
	root_template.Execute(w, nil)
	http.FileServer(http.Dir("templates/root.html"))
}

// /c/:chan:
func (s *serverContext) loadChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if websocket.IsWebSocketUpgrade(r) {
		s.users = map[string]*appUser{"a": &appUser{c: &conn{authenticated: false, websocket: nil}, channels: nil}}
		s.users["a"].c.upgrade(w, r)
		fmt.Printf("%+v", s)
	}
	// Get the complete URL path into an array
	ci := s.getChannelInfo(r.URL.Path, params)
	js_asset, _ := Asset("assets/channel.js")
	css_asset, _ := Asset("assets/channel.css")
	scriptTemplate := &templateAssets{
		JS:  template.JS(js_asset),
		CSS: string(css_asset),
	}

	log := getLog(ci)
	template_data := struct {
		ChannelName string
		ChannelLog  string
		Template    *templateAssets
	}{
		ChannelName: ci.ChannelName,
		ChannelLog:  log,
		Template:    scriptTemplate,
	}
	data, err := Asset("templates/channel.html")
	if err != nil {
		fmt.Println("Error acquiring channel.html asset.")
	}

	channel_template := template.New("channel")
	channel_template, err = channel_template.Parse(string(data))
	if err != nil {
		fmt.Println(err)
	}
	err = channel_template.Execute(w, template_data)
	if err != nil {
		fmt.Println(err)
	}
}

// /c/:chan:
func (sc *serverContext) saveChannel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	accessLog(r)
	// Save the state of the channel to a file
	// TODO: This needs to work with WebSockets so that new messages are pushed
	// to the client. Right now the user has to refresh the page.
	ci := sc.getChannelInfo(r.URL.Path, params)
	s := getSession(w, r)
	fmt.Printf("%+v", sc)
	sc.users["a"].c.websocket.WriteMessage(websocket.TextMessage, []byte("hey"))
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
	sc.loadChannel(w, r, params)
}

// /login
func (s *serverContext) getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	accessLog(r)
	data, err := Asset("templates/login.html")
	if err != nil {
		fmt.Println("Error acquiring root.html asset.")
	}
	login_template := template.New("login")
	login_template.Parse(string(data))
	login_template.Execute(w, nil)
}

// /login
func (s *serverContext) postLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	accessLog(r)
	r.ParseForm()
	if r.Form["username"][0] == "" {
		fmt.Fprint(w, "Please enter a username.")
	} else {
		fmt.Println("username: ", r.PostForm["username"])
		fmt.Println("password: ", r.PostForm["password"])
	}
}
