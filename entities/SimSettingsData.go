package entities

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type SimSettingsData struct {
	NumberOfSims  int
	NumberOfTurns int
	Drops         map[int]int
}

func NewSimSettingsData(req *http.Request) (SimSettingsData, error) {
	drops := make(map[int]int)
	for key, values := range req.Form {
		if strings.Contains(key, "drops") && len(values) > 0 {
			res, err := strconv.Atoi(values[0])
			if err != nil {
				return SimSettingsData{}, err
			}
			intKey, err := strconv.Atoi(strings.TrimSuffix(key, "drops"))
			if err != nil {
				return SimSettingsData{}, err
			}
			drops[intKey] = res
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
		NumberOfSims:  numSims,
		NumberOfTurns: numTurns,
		Drops:         drops,
	}, nil
}

func (s SimSettingsData) GetNumberOfSims() int {
	return s.NumberOfSims
}

func (s SimSettingsData) GetNumberOfTurns() int {
	return s.NumberOfTurns
}

func (s SimSettingsData) GetNumberOfSpells() int {
	spellCount := 0
	for _, value := range s.Drops {
		spellCount += value
	}
	return spellCount
}

func (s SimSettingsData) GetOrderedListOfDrops() []int {
	dropValues := make([]int, 0, len(s.Drops))
	for cost, numSpellsAtCost := range s.Drops {
		if numSpellsAtCost > 0 {
			dropValues = append(dropValues, cost)
		}
	}
	sort.Slice(dropValues, func(i, j int) bool {
		return dropValues[i] < dropValues[j]
	})
	return dropValues
}

func (s SimSettingsData) GetListOfSpells() []BasicSpell {
	spells := make([]BasicSpell, 0, len(s.Drops))
	orderedListOfDrops := s.GetOrderedListOfDrops()
	for _, cost := range orderedListOfDrops {
		for i := 0; i < s.Drops[cost]; i++ {
			spells = append(spells, BasicSpell{cost: cost})
		}
	}
	return spells
}
