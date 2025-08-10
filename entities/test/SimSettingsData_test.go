package entities

import (
	"net/http"
	"testing"

	"github.com/menterline/oncurve/entities"
)

func createTestSimSettingsData() (entities.SimSettingsData, error) {
	req, err := http.NewRequest("POST", "/simulate", nil)
	if err != nil {
		return entities.SimSettingsData{}, err
	}
	req.Form = make(map[string][]string)
	req.Form.Add("numberOfSims", "1000")
	req.Form.Add("numberOfTurns", "10")
	req.Form.Add("1drops", "4")
	req.Form.Add("2drops", "3")
	req.Form.Add("5drops", "1")
	return entities.NewSimSettingsData(req)
}

func TestFactory(t *testing.T) {
	result, err := createTestSimSettingsData()
	if err != nil {
		t.Errorf("Error create test data")
	}
	if result.GetNumberOfSims() != 1000 {
		t.Errorf("Expected numberOfSims to be '1000', got '%d'", result.GetNumberOfSims())
	}
	if result.GetNumberOfTurns() != 10 {
		t.Errorf("Expected numberOfTurns to be '10', got '%d'", result.GetNumberOfTurns())
	}
	if len(result.Drops) != 3 {
		t.Errorf("Expected 3 drops, got %d", len(result.Drops))
	}
	if result.Drops[1] != 4 {
		t.Errorf("Expected 1drops to be '4', got '%d'", result.Drops[1])
	}
	if result.Drops[2] != 3 {
		t.Errorf("Expected 2drops to be '3', got '%d'", result.Drops[2])
	}
	if result.Drops[5] != 1 {
		t.Errorf("Expected 5drops to be '1', got '%d'", result.Drops[5])
	}
}

func TestFactory_ExpectError(t *testing.T) {
	req, err := http.NewRequest("POST", "/simulate", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Form = make(map[string][]string)
	req.Form.Add("THIS IS A BAD STRING", "1000")
	_, err = entities.NewSimSettingsData(req)

	if err == nil {
		t.Error("Expected an error due to bad string in form, but got none")
	}
}

func TestGetNumberOfSpells(t *testing.T) {
	result, err := createTestSimSettingsData()
	if err != nil {
		t.Errorf("Error create test data")
	}
	if result.GetNumberOfSpells() != 8 {
		t.Errorf("Expected number of spells to be 8, got %d", result.GetNumberOfSpells())
	}
}

func TestGetListOfSpells(t *testing.T) {
	result, err := createTestSimSettingsData()
	if err != nil {
		t.Errorf("Error create test data")
	}
	spells := result.GetListOfSpells()
	if len(spells) != 8 {
		t.Errorf("Expected 8 spells, got %d", len(spells))
	}
	for i, spell := range spells {
		if i < 4 && spell.GetManaCost() != 1 {
			t.Errorf("Expected spell %d to have mana cost %d, got %d", i, 1, spell.GetManaCost())
		}
		if i > 3 && i < 7 && spell.GetManaCost() != 2 {
			t.Errorf("Expected spell %d to have mana cost %d, got %d", i, 2, spell.GetManaCost())
		}
		if i == 7 && spell.GetManaCost() != 5 {
			t.Errorf("Expected spell %d to have mana cost %d, got %d", i, 5, spell.GetManaCost())
		}
	}
}
