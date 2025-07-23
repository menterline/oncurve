package services

import (
	"fmt"
	"net/http"
)

type SimulationResult struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

func Simulate(parsedForm *http.Request) SimulationResult {
	formMap := formToMap(parsedForm)
	fmt.Println(formMap)
	return SimulationResult{
		Successes: 400,
		Failures:  100,
	}
}

func formToMap(parsedForm *http.Request) map[string]string {
	result := make(map[string]string)
	for key, values := range parsedForm.Form {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}
	return result
}
