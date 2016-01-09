package main

import (
	"fmt"
	"net/http"

	"github.com/fuzzybear3965/go-chat/static"
	"github.com/julienschmidt/httprouter"
)

func getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET/POST): ", r.Method)
	logintmpl := static.Login
	logintmpl.Execute(w, nil)
}

func postLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET/POST): ", r.Method)
	r.ParseForm()
	if r.Form["username"][0] == "" {
		fmt.Fprint(w, "Please enter a username.")
	} else {
		fmt.Println("username: ", r.PostForm["username"])
		fmt.Println("password: ", r.PostForm["password"])
	}
}
