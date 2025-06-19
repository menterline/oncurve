package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/hello", sayHello)

	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/hello.html"))
	tmpl.Execute(w, nil)
}
