package entities

import (
	"net/http"
	"testing"

	"github.com/menterline/oncurve/entities"
)

func TestFormToMap(t *testing.T) {
	// Create a mock HTTP request with form values
	req, err := http.NewRequest("POST", "/simulate", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Form = make(map[string][]string)
	req.Form.Add("numberOfSims", "1000")
	req.Form.Add("numberOfTurns", "10")
	req.Form.Add("1drops", "4")
	req.Form.Add("2drops", "3")
	req.Form.Add("5drops", "1")
	// Call the FormToMap function
	result := entities.SimSettingsDataFromForm(req)

	// Check if the result contains the expected keys and values
	if result.GetNumberOfSims() != 1000 {
		t.Errorf("Expected numberOfSims to be '1000', got '%d'", result.GetNumberOfSims())
	}
	if result.GetNumberOfTurns() != 10 {
		t.Errorf("Expected numberOfTurns to be '10', got '%d'", result.GetNumberOfTurns())
	}
	if len(result.Drops) != 3 {
		t.Errorf("Expected 3 drops, got %d", len(result.Drops))
	}
	if result.Drops["1drops"] != 4 {
		t.Errorf("Expected 1drops to be '4', got '%d'", result.Drops["1drops"])
	}
	if result.Drops["2drops"] != 3 {
		t.Errorf("Expected 2drops to be '3', got '%d'", result.Drops["2drops"])
	}
	if result.Drops["5drops"] != 1 {
		t.Errorf("Expected 5drops to be '1', got '%d'", result.Drops["5drops"])
	}

}
