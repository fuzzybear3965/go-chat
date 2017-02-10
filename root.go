package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func (s *serverContext) rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET): ", r.Method)
	data, err := Asset("templates/root.html")
	if err != nil {
		fmt.Println("Error acquiring root.html asset.")
	}
	root_template := template.New("root")
	root_template.Parse(string(data))
	root_template.Execute(w, nil)
}
