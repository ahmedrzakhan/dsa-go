package main

import "fmt"

// Graph representation using adjacency lists
type GraphTS map[int][]int

// TC - O(V+E), SC - O(V+E)
func topologicalSort(edges [][]int, n int) []int {
	adjList := make(GraphTS) // Initialize the graph

	// Build the adjacency list
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		adjList[src] = append(adjList[src], dst)
	}

	visited := make(map[int]bool)
	result := []int{}

	// Perform DFS starting from each node
	for i := 1; i <= n; i++ {
		dfsTS(i, adjList, visited, &result)
	}

	// Reverse the result for correct topological order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func dfsTS(src int, adjList GraphTS, visited map[int]bool, result *[]int) {
	if visited[src] {
		return
	}
	visited[src] = true

	for _, neighbor := range adjList[src] {
		dfsTS(neighbor, adjList, visited, result)
	}
	*result = append(*result, src)
}

func mainTS() {
	// Example usage
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}}
	n := 4
	order := topologicalSort(edges, n)
	fmt.Println(order)
}
