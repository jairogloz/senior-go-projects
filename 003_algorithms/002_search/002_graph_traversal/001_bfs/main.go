package main

import (
	"fmt"
	"slices"
)

func main() {

	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 5},
		3: {1},
		4: {1, 5},
		5: {2, 4},
	}

	fmt.Println("BFS Traversal:")
	bfs(graph, 0)
}

func bfs(graph map[int][]int, start int) {
	_, ok := graph[start]
	if !ok {
		return
	}
	// queue: [0]
	queue := []int{start}
	// visited: [0]
	visited := []int{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current) // Process node
		// Get the current node
		node, _ := graph[current]
		for _, neighbor := range node {
			if !slices.Contains(visited, neighbor) {
				queue = append(queue, neighbor)
				visited = append(visited, neighbor)
			}
		}
	}

}
