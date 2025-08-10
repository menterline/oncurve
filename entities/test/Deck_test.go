package entities_test

import (
	"testing"

	"github.com/menterline/oncurve/entities"
)

func TestNewDeck(t *testing.T) {
	testData := entities.SimSettingsData{
		NumberOfSims:  1000,
		NumberOfTurns: 10,
		Drops: map[int]int{
			1: 7,
			2: 6,
			3: 4,
			4: 4,
			5: 1,
			6: 3,
		},
	}
	deck := entities.NewDeck(testData)
	if deck.GetSize() != 60 {
		t.Errorf("Expected deck length to be 60, got %d", deck.GetSize())
	}
	spellCount := 0
	landCount := 0
	for _, card := range deck.GetCards() {
		if card.GetCardType() == entities.Spell {
			spellCount++
		}
		if card.GetCardType() == entities.Land {
			landCount++
		}
	}
	if spellCount != 25 {
		t.Errorf("Expected 25 spells, got %d", spellCount)
	}
	if landCount != 35 {
		t.Errorf("Expected 35 lands, got %d", landCount)
	}

}

func TestDraw(t *testing.T) {
	testCards := [5]entities.Card{}
	testCards[0] = entities.BasicLand{}
	testCards[1] = entities.BasicLand{}
	testCards[2] = entities.NewBasicSpell(2)
	testCards[3] = entities.NewBasicSpell(1)
	testCards[4] = entities.NewBasicSpell(3)
	deck, _ := entities.NewDeckFromCards(testCards[:])
	newCard, _ := deck.Draw()
	if deck.GetSize() != 4 {
		t.Errorf("Expected deck size to be 4 after drawing one card, got %d", deck.GetSize())
	}
	if newCard.GetCardType() != entities.Land {
		t.Errorf("Expected drawn card to be a land, got %d", newCard.GetCardType())
	}
	deck.Draw()
	newCard, _ = deck.Draw()
	if deck.GetSize() != 2 {
		t.Errorf("Expected deck size to be 2 after drawing two more cards, got %d", deck.GetSize())
	}
	if newCard.GetCardType() != entities.Spell {
		t.Errorf("Expected drawn card to be a land, got %d", newCard.GetCardType())
	}
	if newCard.GetManaCost() != 2 {
		t.Errorf("Expected drawn card to be cost 2, got %d", newCard.GetManaCost())
	}
}
