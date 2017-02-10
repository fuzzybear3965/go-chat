package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	//"github.com/julienschmidt/httprouter"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type conn struct {
	authenticated bool
	websocket     *websocket.Conn
}

func (c *conn) upgrade(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		fmt.Println("had an error upgrading")
	}

	ws.WriteMessage(websocket.TextMessage, []byte("Hi, John. I'm a websocket, now."))

	c.websocket = ws
}
