package main

import "fmt"

// TC - O(V+E), SC - O(V+E)
func detectCycle(edges [][]int) bool {
	adjList := make(map[int][]int)
	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		if _, ok := adjList[src]; !ok {
			adjList[src] = append(adjList[src], dst)
		}
	}

	for node := range adjList {
		if !visited[node] {
			if dfsCD(node, adjList, visited, visiting) {
				return true
			}
		}
	}

	return false
}

func dfsCD(node int, adjList map[int][]int, visited, visiting map[int]bool) bool {
	if visiting[node] {
		return true // Cycle detected
	}

	if visited[node] {
		return false // Node has been fully explored and no cycle was found on this path
	}

	visited[node] = true
	visiting[node] = true

	for _, neighbor := range adjList[node] {
		if dfsCD(neighbor, adjList, visited, visiting) {
			return true
		}
	}

	visiting[node] = false // Remove from recursion stack when backtracking
	return false
}

func mainCD() {
	edgesWithCycle := [][]int{{1, 0}, {2, 1}, {0, 2}}
	fmt.Println("Graph with cycle:", detectCycle(edgesWithCycle))

	edgesWithoutCycle := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println("Graph without cycle:", detectCycle(edgesWithoutCycle))
}
