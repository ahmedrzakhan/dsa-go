package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node1, node2, weight int
}

type EdgeHeap []Edge

// Implement heap.Interface for EdgeHeap
func (h EdgeHeap) Len() int {
	return len(h)
}

func (h EdgeHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

func (h EdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	el := (*h)[len(*h)-1] // Get the last element
	*h = (*h)[:len(*h)-1] // Shrink the slice
	return el
}

// TC - O(ELogV), SC - O(E)
func minimumSpanningTree(edges [][]int, N int) ([][]int, int) {
	adjList := make(map[int][]Edge) // Adjacency list representation

	// Build the adjacency list
	for _, edge := range edges {
		N1, N2, weight := edge[0], edge[1], edge[2]
		adjList[N1] = append(adjList[N1], Edge{node1: N1, node2: N2, weight: weight})
		adjList[N2] = append(adjList[N2], Edge{node1: N2, node2: N1, weight: weight})
	}

	minHeap := &EdgeHeap{} // Initialize min-heap
	heap.Init(minHeap)

	visited := make(map[int]bool)
	mst := [][]int{}
	totalWeight := 0

	// Start with node 1
	visited[1] = true

	// Push all edges connected to node 1 onto the heap
	for _, e := range adjList[1] {
		heap.Push(minHeap, e)
	}

	for len(*minHeap) > 0 && len(visited) < N {
		edge := heap.Pop(minHeap).(Edge)
		N1, N2, weight := edge.node1, edge.node2, edge.weight

		// Check if the neighboring node is already visited
		if visited[N2] {
			continue
		}

		// Add the edge to the MST
		mst = append(mst, []int{N1, N2})
		totalWeight += weight
		visited[N2] = true

		// Push all edges connected to the newly visited node onto the heap
		for _, neighborEdge := range adjList[N2] {
			if !visited[neighborEdge.node2] {
				heap.Push(minHeap, neighborEdge)
			}
		}
	}

	return mst, totalWeight
}

func mainP() {
	// Example usage
	edges := [][]int{{1, 2, 3}, {1, 3, 1}, {2, 3, 4}}
	n := 3
	mst, totalWeight := minimumSpanningTree(edges, n)
	fmt.Println("mst", mst)
	fmt.Println("totalWeight", totalWeight)
}

/**
    1 ---- 3
   / \
  3   1
 /
2 ---- 4
*/
