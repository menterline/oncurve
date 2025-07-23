package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/menterline/oncurve/services"
)

func SimulateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	result := services.Simulate(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
