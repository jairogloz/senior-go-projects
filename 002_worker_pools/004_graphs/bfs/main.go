package main

import (
	"fmt"
)

// BFS to find the shortest path in an unweighted graph
func BFSShortestPath(graph map[string][]string, start, target string) int {
	// Queue to hold nodes and their distances from the start
	queue := []struct {
		node     string
		distance int
	}{{start, 0}}

	// Visited set to track visited nodes
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		// Dequeue the first element
		current := queue[0]
		queue = queue[1:]

		// If we reached the target, return the distance
		if current.node == target {
			return current.distance
		}

		// Enqueue all unvisited neighbors
		for _, neighbor := range graph[current.node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, struct {
					node     string
					distance int
				}{neighbor, current.distance + 1})
			}
		}
	}

	// If target is not reachable
	return -1
}

func main() {
	// Define the graph as an adjacency list
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}

	// Test BFSShortestPath
	start := "F"
	target := "A"
	result := BFSShortestPath(graph, start, target)

	if result != -1 {
		fmt.Printf("The shortest path from %s to %s is %d steps.\n", start, target, result)
	} else {
		fmt.Printf("There is no path from %s to %s.\n", start, target)
	}
}
