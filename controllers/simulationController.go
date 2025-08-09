package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/menterline/oncurve/entities"
	"github.com/menterline/oncurve/services"
)

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
	json.NewEncoder(w).Encode(result)
}
