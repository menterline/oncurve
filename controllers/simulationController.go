package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/menterline/oncurve/services"
)

func SimulateHandler(w http.ResponseWriter, r *http.Request) {
	result := services.Simulate()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
