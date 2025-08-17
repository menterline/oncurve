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

/*
DFS to see if there is a path that succeeds
*/
func RunSimulation(numTurns int, deck entities.Deck) bool {

	/*
		 while currentHandNumber < numTurns {
			draw
				if no lands, return false
				if no spells, return false
	*/
	return false
}

func onCurveDFS(currentHandNumber int, numberOfTurns int, deck entities.Deck, boardState entities.BoardState) bool {
	return true
}
