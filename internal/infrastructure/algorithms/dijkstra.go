package algorithms

import (
	"math"

	"shortest-path-api/internal/domain"
)

func Dijkstra(graph domain.Graph, start string) (map[string]float64, map[string]string) {
	dist := make(map[string]float64)
	prev := make(map[string]string)
	unvisited := make(map[string]bool)

	for node := range graph {
		dist[node] = math.Inf(1)
		unvisited[node] = true
	}
	dist[start] = 0

	for len(unvisited) > 0 {
		var minNode string
		minDist := math.Inf(1)
		for node := range unvisited {
			if dist[node] < minDist {
				minDist = dist[node]
				minNode = node
			}
		}

		if minNode == "" {
			break
		}
		delete(unvisited, minNode)

		for neighbor, weight := range graph[minNode] {
			alt := dist[minNode] + weight
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = minNode
			}
		}
	}

	return dist, prev
}

func BuildPath(prev map[string]string, target string) []string {
	var path []string
	for target != "" {
		path = append([]string{target}, path...)
		target = prev[target]
	}
	return path
}
