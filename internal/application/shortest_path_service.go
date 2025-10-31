package application

import (
	"math"

	"shortest-path-api/internal/domain"
	"shortest-path-api/internal/infrastructure/algorithms"
)

type ShortestPathService struct{}

func NewShortestPathService() *ShortestPathService {
	return &ShortestPathService{}
}

func (s *ShortestPathService) FindClosestDepot(input domain.Input) (*domain.Output, error) {
	graph := input.Graph
	accident := input.AccidentLocation

	var bestDepot string
	var bestDistance = math.Inf(1)
	var bestPath []string

	for _, depot := range input.Depots {
		dist, prev := algorithms.Dijkstra(graph, depot)
		if d, ok := dist[accident]; ok && d < bestDistance {
			bestDistance = d
			bestDepot = depot
			bestPath = algorithms.BuildPath(prev, accident)
		}
	}

	if math.IsInf(bestDistance, 1) {
		return nil, domain.ErrNoPathFound
	}

	return &domain.Output{
		FromDepot: bestDepot,
		To:        accident,
		Path:      bestPath,
		Distance:  bestDistance,
	}, nil
}
