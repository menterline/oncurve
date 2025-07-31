package services

import (
	"fmt"
	"net/http"

	"github.com/menterline/oncurve/entities"
)

type SimulationResult struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

func Simulate(parsedForm *http.Request) SimulationResult {
	simData := entities.SimSettingsDataFromForm(parsedForm)
	fmt.Println(simData)
	return SimulationResult{
		Successes: 400,
		Failures:  100,
	}
}
