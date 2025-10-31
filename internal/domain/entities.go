package domain

type Graph map[string]map[string]float64

type Input struct {
	AccidentLocation string   `json:"accidentLocation"`
	Depots           []string `json:"depots"`
	Graph            Graph    `json:"graph"`
}

type Output struct {
	FromDepot string   `json:"fromDepot"`
	To        string   `json:"to"`
	Path      []string `json:"path"`
	Distance  float64  `json:"distance"`
}