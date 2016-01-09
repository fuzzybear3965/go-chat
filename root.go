package main

import (
	"fmt"
	"net/http"

	"github.com/fuzzybear3965/go-chat/static"
	"github.com/julienschmidt/httprouter"
)

func loadRoot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Method (GET): ", r.Method)
	roottmpl := static.Root
	roottmpl.Execute(w, nil)
	//t := template.New()
	//t, err := t.ParseFiles("root.gtpl")
	//if err != nil {
	//	fmt.Fprintln(w, "Could not read root template.")
	//	log.Fatal(err)
	//}

	//t.Execute(w)
}
