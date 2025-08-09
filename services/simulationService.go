package services

import (
	"fmt"

	"github.com/menterline/oncurve/entities"
)

type SimulationResult struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

func RunAllSimulations(simData entities.SimSettingsData) SimulationResult {
	successes := 0
	failures := 0
	for i := 0; i < simData.GetNumberOfSims(); i++ {
		result := simData.Simulate(RunSimulation)
		if result {
			successes++
		} else {
			failures++
		}
	}
	fmt.Println(simData)
	return SimulationResult{
		Successes: successes,
		Failures:  failures,
	}
}

func RunSimulation() {

}
