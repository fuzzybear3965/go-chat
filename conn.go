package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	//"github.com/julienschmidt/httprouter"
	"net/http"
)

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

	c.websocket = ws
}
