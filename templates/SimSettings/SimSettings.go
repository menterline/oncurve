package simsettings

import (
	"html/template"
	"net/http"
)

type ConvertedManaCosts struct {
	CMCs []int
}

func ServeSimSettings(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/SimSettings/index.html"))
	data := ConvertedManaCosts{
		CMCs: []int{1, 2, 3, 4 },
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
