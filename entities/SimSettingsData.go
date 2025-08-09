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

func NewSimSettingsData(req *http.Request) (SimSettingsData, error) {
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
		return SimSettingsData{}, err
	}
	numTurns, err := strconv.Atoi(req.Form.Get("numberOfTurns"))
	if err != nil {
		return SimSettingsData{}, err
	}
	return SimSettingsData{
		numberOfSims:  numSims,
		numberOfTurns: numTurns,
		Drops:         drops,
	}, nil
}

func (s SimSettingsData) Simulate(runSim func()) bool {
	return true
}
