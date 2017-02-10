package main

import (
	//"github.com/gorilla/websocket"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//type conn struct {
//username      string
//authenticated bool
//websocket     *websocket.Conn
//channels      []string
//}

//func (c *conn) upgrade(w http.ResponseWriter, r *http.Request) {
//ws, err := upgrader.Upgrade(w, r, nil)

//if err != nil {
//fmt.Println(err)
//fmt.Println("had an error upgrading")
//}

//c.websocket = ws
//}

//func (ctx *context) ServeHTTP(w http.ResponseWriter, r *http.Request) {

//}

func (ctx *context) customHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(ctx.a, "okay, this works.")
}

type context struct {
	a string
}
