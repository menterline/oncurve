package main

import (
	"html/template"
	"net/http"

	simControllers "github.com/menterline/oncurve/controllers"
	simsettings "github.com/menterline/oncurve/templates/SimSettings"
)

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/SimSettings", simsettings.ServeSimSettings)
	http.HandleFunc("/Simulate", simControllers.SimulateHandler)
	http.HandleFunc("/AddNewMana", simsettings.AddNewManaHandler)
	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
