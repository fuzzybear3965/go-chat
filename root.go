package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func loadRoot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "root.gtpl")
	//t := template.New()
	//t, err := t.ParseFiles("root.gtpl")
	//if err != nil {
	//	fmt.Fprintln(w, "Could not read root template.")
	//	log.Fatal(err)
	//}

	//t.Execute(w)
}
