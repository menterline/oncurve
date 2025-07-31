package services

import (
	"fmt"
	"net/http"
	"strconv"
)

type SimulationResult struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

func Simulate(parsedForm *http.Request) SimulationResult {
	formMap := FormToMap(parsedForm)
	fmt.Println(formMap)
	return SimulationResult{
		Successes: 400,
		Failures:  100,
	}
}

func FormToMap(parsedForm *http.Request) map[string]int {
	result := make(map[string]int)
	for key, values := range parsedForm.Form {
		if len(values) > 0 {
			res, err := strconv.Atoi(values[0])
			if err != nil {
				fmt.Printf("Error converting %s to int: %v\n", values[0], err)
			}
			result[key] = res
		}
	}
	numSims, err := strconv.Atoi(parsedForm.Form.Get("numberOfSims"))
	if err != nil {
		fmt.Printf("Error converting numberOfSims to int: %v\n", err)
	}
	numTurns, err := strconv.Atoi(parsedForm.Form.Get("numberOfTurns"))
	if err != nil {
		fmt.Printf("Error converting numberOfTurns to int: %v\n", err)
	}
	result["numberOfSims"] = numSims
	result["numberOfTurns"] = numTurns
	return result
}
