package main

import "fmt"

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
	dfs(graph, 0, make(map[int]bool))
}

func dfs(graph map[int][]int, start int, visited map[int]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println(start) // Process node

	for _, neighbor := range graph[start] {
		dfs(graph, neighbor, visited)
	}
}
