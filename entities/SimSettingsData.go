package entities

import (
	"log"
	"net/http"
	"strconv"
)

type SimSettingsData struct {
	NumberOfSims  int
	NumberOfTurns int
	Drops         map[string]int
}

func SimSettingsDataFromForm(req *http.Request) SimSettingsData {
	drops := make(map[string]int)
	for key, values := range req.Form {
		if len(values) > 0 {
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
		NumberOfSims:  numSims,
		NumberOfTurns: numTurns,
		Drops:         drops,
	}
}
