package test

import (
	"testing"

	"github.com/menterline/oncurve/entities"
	"github.com/menterline/oncurve/services"
)

func TestRunSimulationDeckAllLands_ExpectFalse(t *testing.T) {
	cards := make([]entities.Card, entities.MAX_DECK_SIZE)
	for i := 0; i < entities.MAX_DECK_SIZE; i++ {
		cards[i] = entities.BasicLand{}
	}
	deck := entities.Deck{Cards: cards}
	result := services.RunSimulation(1, deck)
	if result != false {
		t.Errorf("Expected false, got %t", result)
	}
}

func TestRunSimulationDeckAllZeroCostSpells_ExpectTrue(t *testing.T) {
	cards := make([]entities.Card, entities.MAX_DECK_SIZE)
	for i := 0; i < entities.MAX_DECK_SIZE; i++ {
		cards[i] = entities.BasicSpell{}
	}
	deck := entities.Deck{Cards: cards}
	result := services.RunSimulation(1, deck)
	if result != true {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestRunSimulationDeckAllNonZeroCostSpells_ExpectFalse(t *testing.T) {
	cards := make([]entities.Card, entities.MAX_DECK_SIZE)
	for i := 0; i < entities.MAX_DECK_SIZE; i++ {
		cards[i] = entities.NewBasicSpell(1)
	}
	deck := entities.Deck{Cards: cards}
	result := services.RunSimulation(1, deck)
	if result != false {
		t.Errorf("Expected false, got %t", result)
	}
}

func TestRunSimulationDeckAllBigSpells_ExpectFalse(t *testing.T) {
	cards := make([]entities.Card, 10)
	for i := 0; i < 5; i++ {
		cards[i] = entities.BasicLand{}
	}
	for i := 5; i < 10; i++ {
		cards[i] = entities.NewBasicSpell(5)
	}
	deck := entities.Deck{Cards: cards}
	result := services.RunSimulation(2, deck)
	if result != false {
		t.Errorf("Expected false, got %t", result)
	}
}

func TestRunSimulationDeckAllLandsAndOneDrops_ExpectTrue(t *testing.T) {
	cards := make([]entities.Card, 10)
	for i := 0; i < 5; i++ {
		cards[i] = entities.BasicLand{}
	}
	for i := 5; i < 10; i++ {
		cards[i] = entities.NewBasicSpell(1)
	}
	deck := entities.Deck{Cards: cards}
	result := services.RunSimulation(3, deck)
	if result != true {
		t.Errorf("Expected true, got %t", result)
	}
}

func TestRunSimulationHitDrops_ExpectTrue(t *testing.T) {
	cards := make([]entities.Card, 10)
	for i := 0; i < 6; i++ {
		cards[i] = entities.BasicLand{}
	}
	cards[6] = entities.NewBasicSpell(1)
	cards[7] = entities.NewBasicSpell(2)
	cards[8] = entities.NewBasicSpell(3)
	cards[9] = entities.NewBasicSpell(4)
	deck := entities.Deck{Cards: cards}
	result := services.RunSimulation(4, deck)
	if result != true {
		t.Errorf("Expected true, got %t", result)
	}
}
