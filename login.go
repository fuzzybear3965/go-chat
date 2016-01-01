package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"text/template"
)

func getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET/POST): ", r.Method)
	tmpl, err := template.ParseFiles("login.gtpl")
	if err != nil {
		fmt.Println("Could not load login template.")
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
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
