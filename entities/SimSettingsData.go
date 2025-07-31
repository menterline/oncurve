package entities

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SimSettingsData struct {
	numberOfSims  int
	numberOfTurns int
	Drops         map[string]int
}

func (s SimSettingsData) GetNumberOfSims() int {
	return s.numberOfSims
}

func (s SimSettingsData) GetNumberOfTurns() int {
	return s.numberOfTurns
}

func SimSettingsDataFromForm(req *http.Request) SimSettingsData {
	drops := make(map[string]int)
	for key, values := range req.Form {
		if strings.Contains(key, "drops") && len(values) > 0 {
			res, err := strconv.Atoi(values[0])
			if err != nil {
				log.Fatalf("Error converting %s to int: %v\n", values[0], err)
			}
			drops[key] = res
		}
	}
	numSims, err := strconv.Atoi(req.Form.Get("numberOfSims"))
	if err != nil {
		log.Fatalf("Error converting numberOfSims to int: %v\n", err)
	}
	numTurns, err := strconv.Atoi(req.Form.Get("numberOfTurns"))
	if err != nil {
		log.Fatalf("Error converting numberOfTurns to int: %v\n", err)
	}
	return SimSettingsData{
		numberOfSims:  numSims,
		numberOfTurns: numTurns,
		Drops:         drops,
	}
}
