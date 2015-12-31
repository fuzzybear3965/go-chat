package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Page struct {
	Title string
	Body  []byte
}

func open(title string) (*Page, error) {
	filename := title + ".txt"
	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: filebytes}, err
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//	for _, val := range p.Body {
	//		fmt.Println(val)
	//	}
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func defaultPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Welcome to the %q page.", path)
	outputstring := fmt.Sprintf("Welcome to the %q page.", path)
	page := &Page{Title: path, Body: []byte(outputstring)}
	err := page.save()
	if err != nil {
		log.Fatal(err)
	}
}

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

func main() {
	router := httprouter.New()
	router.GET("/", defaultPage)
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)

	log.Fatal(http.ListenAndServe(":8080", router))
}

/****** Me implementing the ServeHTTP method such as to use
http.Handle(PATH string, SOMEFUNC function)
type fooHandler func(http.ResponseWriter, *http.Request)

var fooFunction fooHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, foo")
}

func (f fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
*****/
