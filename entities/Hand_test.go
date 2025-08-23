package entities

import (
	"testing"
)

func TestHand_PlayCard(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			BasicLand{},
			NewBasicSpell(2),
			NewBasicSpell(1),
		},
	}

	boardState := &BoardState{}

	// Test playing a land
	hand.PlayCard(hand.Cards[0], boardState)
	if len(boardState.Lands) != 1 {
		t.Error("Expected to play land, but failed")
	}
	if len(hand.Cards) != 2 {
		t.Error("Expected hand to have 2 cards after playing land, got", len(hand.Cards))
	}

	// Test playing a land
	hand.PlayCard(hand.Cards[1], boardState)
	if len(boardState.Lands) != 1 {
		t.Error("Expected to play spell, but failed")
	}
	if len(hand.Cards) != 1 {
		t.Error("Expected hand to have 1 cards after playing land, got", len(hand.Cards))
	}
}
