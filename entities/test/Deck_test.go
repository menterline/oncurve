package entities

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
}

func TestDraw(t *testing.T) {
	testCards := [5]entities.Card{}
	testCards[0] = entities.BasicLand{}
	testCards[1] = entities.BasicLand{}
	testCards[2] = entities.NewBasicSpell(2)
	testCards[3] = entities.NewBasicSpell(1)
	testCards[4] = entities.NewBasicSpell(3)

}
