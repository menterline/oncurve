package services

type SimulationResult struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

func Simulate() SimulationResult {
	return SimulationResult{
		Successes: 400,
		Failures:  100,
	}
}
