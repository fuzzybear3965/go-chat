package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func loadRoot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET): ", r.Method)
	data, err := Asset("static/root.html")
	if err != nil {
		fmt.Println("Error acquiring root.html asset.")
	}
	root_template := template.New("root")
	root_template.Parse(string(data))
	root_template.Execute(w, nil)
	//t := template.New()
	//t, err := t.ParseFiles("root.gtpl")
	//if err != nil {
	//	fmt.Fprintln(w, "Could not read root template.")
	//	log.Fatal(err)
	//}

	//t.Execute(w)
}
