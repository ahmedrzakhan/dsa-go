package main

import (
	"fmt"
)

// GraphNode used for adjacency list
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: make([]*GraphNode, 0),
	}
}

func BuildAdjList() map[string][]string {
	adjList := make(map[string][]string)
	// Given directed edges, build an adjacency list
	edges := [][]string{{"A", "B"}, {"B", "C"}, {"B", "E"}, {"C", "E"}, {"E", "D"}}

	adjList["A"] = make([]string, 0)
	adjList["B"] = make([]string, 0)

	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		if _, ok := adjList[src]; !ok {
			adjList[src] = make([]string, 0)
		}
		if _, ok := adjList[dst]; !ok {
			adjList[dst] = make([]string, 0)
		}
		adjList[src] = append(adjList[src], dst)
	}

	return adjList
}

// TODO: iterative approach using stack
// TC - O(V+E), SC - O(V)
// Count paths (backtracking)
func DFS(node, target string, adjList map[string][]string, visited map[string]bool) int {
	if node == target {
		return 1
	}
	if visited[node] {
		return 0
	}

	visited[node] = true
	count := 0
	for _, neighbor := range adjList[node] {
		count += DFS(neighbor, target, adjList, visited)
	}
	visited[node] = false // Reset the visited status after exploring all paths through this node

	return count
}

// TC - O(V+E), SC - O(V)
// Shortest path from node to target
func BFS(node, target string, adjList map[string][]string) int {
	length := 0
	visited := make(map[string]bool)
	queue := []string{node}
	visited[node] = true

	for len(queue) != 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			curr := queue[0]
			queue = queue[1:] // Dequeue
			if curr == target {
				return length
			}
			for _, neighbor := range adjList[curr] {
				if _, ok := visited[neighbor]; !ok {
					visited[neighbor] = true
					queue = append(queue, neighbor) // Enqueue
				}
			}
		}
		length++
	}
	return length
}

func mainAdjList() {
	adjList := BuildAdjList()

	visited := make(map[string]bool)
	fmt.Println("DFS A to D:", DFS("A", "D", adjList, visited))

	fmt.Println("BFS A to D:", BFS("A", "D", adjList))
}
