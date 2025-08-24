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
	startingHandCards, err := deck.DrawMultiple(7)
	if err != nil {
		return false
	}
	startingHand := entities.Hand{Cards: startingHandCards}
	return onCurveDFS(numTurns, startingHand, deck, entities.BoardState{})
}

func onCurveDFS(numberOfTurns int, hand entities.Hand, deck entities.Deck, boardState entities.BoardState) bool {
	if boardState.TurnIndex >= numberOfTurns {
		return true
	}
	plays := entities.AvailablePlays(hand, boardState)
	if len(plays) == 0 && boardState.GetNumUntappedLands() != 0 {
		return false
	}
	for _, play := range plays {
		hand.PlayCard(play, &boardState)
		if boardState.GetNumUntappedLands() == 0 {
			// Used up all lands, so start next turn
			boardState.NextTurn()
			newCard, err := deck.Draw()
			if err != nil {
				return false
			}
			hand.Cards = append(hand.Cards, newCard)
			if onCurveDFS(numberOfTurns, hand, deck, boardState) {
				return true
			}
		} else {
			// still have lands left to use
			if onCurveDFS(numberOfTurns, hand, deck, boardState) {
				return true
			}
		}
	}
	return false
}
