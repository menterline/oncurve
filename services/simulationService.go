package services

import (
	"github.com/menterline/oncurve/entities"
)

type SimulationResult struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

func RunAllSimulations(simData entities.SimSettingsData) SimulationResult {
	successes := 0
	failures := 0
	for i := 0; i < simData.NumberOfSims; i++ {
		deck := entities.NewDeck(simData)
		deck.Shuffle()
		result := RunSimulation(simData.NumberOfTurns, deck)
		if result {
			successes++
		} else {
			failures++
		}
	}
	return SimulationResult{
		Successes: successes,
		Failures:  failures,
	}
}

// TODO
func RunSimulation(numHands int, deck entities.Deck) bool {
	return true
}
