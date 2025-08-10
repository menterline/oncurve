package entities

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	testData := SimSettingsData{
		numberOfSims:  1000,
		numberOfTurns: 10,
		drops: map[int]int{
			1: 7,
			2: 6,
			3: 4,
			4: 4,
			5: 1,
			6: 3,
		},
	}
	deck := NewDeck(testData)
	if deck.GetSize() != 60 {
		t.Errorf("Expected deck length to be 60, got %d", deck.GetSize())
	}
	spellCount := 0
	landCount := 0
	for _, card := range deck.GetCards() {
		if card.GetCardType() == Spell {
			spellCount++
		}
		if card.GetCardType() == Land {
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
	testCards := [5]Card{}
	testCards[0] = BasicLand{}
	testCards[1] = BasicLand{}
	testCards[2] = NewBasicSpell(2)
	testCards[3] = NewBasicSpell(1)
	testCards[4] = NewBasicSpell(4)
	deck := Deck{cards: testCards[:]}
	newCard, _ := deck.Draw()
	if deck.GetSize() != 4 {
		t.Errorf("Expected deck size to be 4 after drawing one card, got %d", deck.GetSize())
	}
	if newCard.GetCardType() != Land {
		t.Errorf("Expected drawn card to be a land, got %d", newCard.GetCardType())
	}
	deck.Draw()
	newCard, _ = deck.Draw()
	if deck.GetSize() != 2 {
		t.Errorf("Expected deck size to be 2 after drawing two more cards, got %d", deck.GetSize())
	}
	if newCard.GetCardType() != Spell {
		t.Errorf("Expected drawn card to be a land, got %d", newCard.GetCardType())
	}
	if newCard.GetManaCost() != 2 {
		t.Errorf("Expected drawn card to be cost 2, got %d", newCard.GetManaCost())
	}
}
