package entities

import (
	"testing"

	"github.com/menterline/oncurve-service/entities"
)

func TestLandFactory(t *testing.T) {
	manaTypes := []entities.Mana{entities.Red, entities.Red, entities.Blue}
	land := entities.LandFactory(manaTypes)
	got := land.GetManaTypes()
	numRed, hasRed := got[entities.Red]
	numBlue, hasBlue := got[entities.Blue]
	if !hasRed || !hasBlue {
		t.Errorf("Expected land to have Red and Blue mana types, got %v", got)
	}
	if numRed != 2 || numBlue != 1 {
		t.Errorf("Expected 2 Red and 1 Blue mana, got %d Red and %d Blue", numRed, numBlue)
	}
}
