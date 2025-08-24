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

func TestAvailablePlays_HasLandOnTurn1_ExpectJustLandPlay(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			BasicLand{},
			NewBasicSpell(2),
			NewBasicSpell(1),
		},
	}

	boardState := BoardState{
		Lands:     []BasicLand{},
		TurnIndex: 0,
	}
	available := AvailablePlays(hand, boardState)
	if len(available) != 1 {
		t.Errorf("Expected 1 available plays, got %d", len(available))
	}
}

func TestAvailablePlays_HasAlreadyPlayedLandOnTurn1_ExpectNoPlays(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			BasicLand{},
			NewBasicSpell(2),
			NewBasicSpell(8),
		},
	}

	boardState := BoardState{
		Lands:     []BasicLand{{}},
		TurnIndex: 0,
	}
	available := AvailablePlays(hand, boardState)
	if len(available) != 0 {
		t.Errorf("Expected 0 available plays, got %d", len(available))
	}
}

func TestAvailablePlays_HasSpellAfterPlayingLand_ExpectOnePlay(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			BasicLand{},
			NewBasicSpell(2),
			NewBasicSpell(1),
		},
	}

	boardState := BoardState{
		Lands:     []BasicLand{{}},
		TurnIndex: 0,
	}
	available := AvailablePlays(hand, boardState)
	if len(available) != 1 {
		t.Errorf("Expected 1 available plays, got %d", len(available))
	}
}

func TestAvailablePlays_HasMultipleLands_ExpectThreePlays(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			BasicLand{},
			BasicLand{},
			BasicLand{},
			NewBasicSpell(2),
			NewBasicSpell(1),
		},
	}

	boardState := BoardState{
		Lands:     []BasicLand{},
		TurnIndex: 0,
	}
	available := AvailablePlays(hand, boardState)
	if len(available) != 3 {
		t.Errorf("Expected 3 available plays, got %d", len(available))
	}
}
