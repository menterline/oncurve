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

func TestHasMana(t *testing.T) {
	manaMap := map[entities.Mana]int{
		entities.Red:   2,
		entities.Blue:  1,
		entities.Green: 3,
	}
	manaPool := entities.ManaPool{AvailableMana: manaMap}

	if !manaPool.HasMana([]entities.Mana{entities.Red, entities.Blue}) {
		t.Errorf("Expected mana pool to have Red and Blue mana")
	}
	if manaPool.HasMana([]entities.Mana{entities.Blue, entities.Blue}) {
		t.Errorf("Expected mana pool to not have enough Blue mana")
	}
}

func TestRemoveFromPoolValid(t *testing.T) {
	manaMap := map[entities.Mana]int{
		entities.Red:   2,
		entities.Blue:  1,
		entities.Green: 3,
	}
	manaPool := entities.ManaPool{AvailableMana: manaMap}
	err1 := manaPool.RemoveFromPool(entities.Blue)
	if err1 != nil {
		t.Errorf("Expected no error, got %v", err1)
	}
	err2 := manaPool.RemoveFromPool(entities.Green)
	if err2 != nil {
		t.Errorf("Expected no error, got %v", err2)
	}
	got := manaPool.GetMana()
	if got[entities.Red] != 2 || got[entities.Blue] != 0 || got[entities.Green] != 2 {
		t.Errorf("Expected mana pool to have 2 Red, 0 Blue, and 3 Green mana, got %v", got)
	}
}

func TestCastSpellValid(t *testing.T) {

	manaMap := map[entities.Mana]int{
		entities.Red:   2,
		entities.Blue:  3,
		entities.Green: 1,
	}
	testManaPool := entities.ManaPool{AvailableMana: manaMap}
	testSpell := entities.Spell{
		Cost:     []entities.Mana{entities.Blue, entities.Blue, entities.Green},
		CardType: entities.Creature}
	testManaPool.CastSpell(testSpell)
	got := testManaPool.GetMana()
	if got[entities.Red] != 2 || got[entities.Blue] != 1 || got[entities.Green] != 0 {
		t.Errorf("Expected mana pool to have 2 Red, 1 Blue, and 0 Green mana, got %v", got)
	}

}
