package main

import (
	"html/template"
	"net/http"
)

type Page interface {
	handler(w http.ResponseWriter, r *http.Request)
}

type StaticPage struct {
	Title    string
	Body     []byte
	Template string
}

type MaxPage struct {
	Title    string
	Body     []byte
	Template string
}

var templates = template.Must(template.ParseFiles("base.html"))

func main() {
	mainPage := StaticPage{Title: "Max is awesome", Body: []byte("Some cool text"), Template: "base.html"}
	secondPage := MaxPage{Title: "Max is awesome2", Body: []byte("Some cool text again because why not"), Template: "base.html"}
	http.HandleFunc("/", makeHandler(mainPage))
	http.HandleFunc("/max", makeHandler(secondPage))
	http.ListenAndServe(":8080", nil)
}

func (s MaxPage) handler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, s.Template, s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s StaticPage) handler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, s.Template, s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(p Page) func(w http.ResponseWriter, r *http.Request) {
	return p.handler
}
