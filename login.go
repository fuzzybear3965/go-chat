package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET/POST): ", r.Method)
	data, err := Asset("templates/login.html")
	if err != nil {
		fmt.Println("Error acquiring root.html asset.")
	}
	login_template := template.New("login")
	login_template.Parse(string(data))
	login_template.Execute(w, nil)
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
