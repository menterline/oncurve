package entities

import (
	"testing"

	"github.com/menterline/oncurve-service/entities"
)

func TestGetManaPool(t *testing.T) {
	manaMap := map[entities.Mana]int{
		entities.Red:   2,
		entities.Blue:  1,
		entities.Green: 3,
	}
	manaPool := entities.ManaPool{AvailableMana: manaMap}
	got := manaPool.GetMana()
	if got[entities.Red] != 2 || got[entities.Blue] != 1 || got[entities.Green] != 3 {
		t.Errorf("Expected mana pool to have 2 Red, 1 Blue, and 3 Green mana, got %v", got)
	}
}

func TestAddToPool(t *testing.T) {
	manaPool := entities.ManaPool{}
	manaPool.AddtoPool(entities.Red, 2)
	manaPool.AddtoPool(entities.Blue, 1)

	got := manaPool.GetMana()
	if got[entities.Red] != 2 || got[entities.Blue] != 1 {
		t.Errorf("Expected mana pool to have 2 Red and 1 Blue mana, got %v", got)
	}	
}
