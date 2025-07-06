package main

import (
	"html/template"
	"net/http"

	deckbuilder "github.com/menterline/oncurve/templates/DeckBuilder"
	simsettings "github.com/menterline/oncurve/templates/SimSettings"
)

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/SimSettings", simsettings.ServeSimSettings)
	http.HandleFunc("/DeckBuilder", deckbuilder.ServeDeckBuilder)
	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
