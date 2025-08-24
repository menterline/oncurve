package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/menterline/oncurve/entities"
	"github.com/menterline/oncurve/services"
)

type SimResults struct {
	TotalSims         int
	NumTurns          int
	Successes         int
	Failures          int
	OnCurvePercentage float64
}

// TODO change this so that the controller parses the request.  service shouldn't know about http
func SimulateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	simData, err := entities.NewSimSettingsData(r)
	if err != nil {
		http.Error(w, "Error parsing Sim Settings Data", http.StatusBadRequest)
		log.Printf("Error creating SimSettingsData: %v", err)
		return
	}
	result := services.RunAllSimulations(simData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("templates/Results/index.html"))
	data := SimResults{
		TotalSims:         simData.NumberOfSims,
		NumTurns:          simData.NumberOfTurns,
		Successes:         result.Successes,
		Failures:          result.Failures,
		OnCurvePercentage: float64(result.Successes) / float64(simData.NumberOfSims) * 100,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO take results and pass to Results/index.html to make look pretty
}
