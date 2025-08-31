package simsettings

import (
	"html/template"
	"net/http"
	"strconv"
)

type ConvertedManaCosts struct {
	CMCs []int
}

func ServeSimSettings(w http.ResponseWriter, r *http.Request) {
	data := ConvertedManaCosts{
		CMCs: generateCMCs(4),
	}
	tmpl := template.Must(template.ParseFiles("templates/SimSettings/index.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddNewManaHandler(w http.ResponseWriter, r *http.Request) {
	currentMax, err := strconv.Atoi(r.URL.Query().Get("currentMax"))
	if err != nil {
		http.Error(w, "Invalid currentMax parameter", http.StatusBadRequest)
		return
	}
	data := ConvertedManaCosts{
		CMCs: generateCMCs((currentMax) + 1),
	}
	tmpl := template.Must(template.ParseFiles("templates/SimSettings/index.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateCMCs(max int) []int {
	cmcs := make([]int, max)
	for i := 0; i < max; i++ {
		cmcs[i] = i + 1
	}
	return cmcs
}
