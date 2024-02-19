package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a vertex in the graph
type NodeDJ struct {
	name     string
	edges    map[string]int // Neighbor nodes and edge weights
	distance int
	visited  bool
}

// Graph represents a weighted graph
type Graph struct {
	nodes map[string]*NodeDJ
}

// An Item for the priority queue
type Item struct {
	node     *NodeDJ
	priority int // Distance
}

// A PriorityQueue implements heap.Interface for Items
type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(name string) {
	g.nodes[name] = &NodeDJ{name: name, edges: make(map[string]int)}
}

// AddEdge adds a weighted edge between two nodes
func (g *Graph) AddEdge(from, to string, weight int) {
	g.nodes[from].edges[to] = weight
}

// Dijkstra finds the shortest path from a source node to a destination node
func (g *Graph) Dijkstra(source, destination string) int {
	// Initialization
	for _, node := range g.nodes {
		node.distance = math.MaxInt32
		node.visited = false
	}
	g.nodes[source].distance = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: g.nodes[source], priority: 0})

	// Main loop
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		curr := item.node

		if curr.visited {
			continue
		}

		curr.visited = true
		if curr.name == destination {
			return curr.distance
		}

		for neighbor, weight := range curr.edges {
			distance := curr.distance + weight
			if distance < g.nodes[neighbor].distance {
				g.nodes[neighbor].distance = distance
				heap.Push(&pq, &Item{node: g.nodes[neighbor], priority: distance})
			}
		}
	}

	return -1 // Destination not found
}

func mainD() {
	graph := Graph{nodes: make(map[string]*NodeDJ)}

	// Sample graph
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	graph.AddEdge("A", "B", 2)
	graph.AddEdge("A", "C", 5)
	graph.AddEdge("B", "C", 1)
	graph.AddEdge("B", "D", 6)
	graph.AddEdge("C", "D", 2)
	graph.AddEdge("C", "E", 3)
	graph.AddEdge("D", "E", 1)

	distance := graph.Dijkstra("A", "E")
	if distance != -1 {
		fmt.Println("Shortest distance from A to E:", distance)
	} else {
		fmt.Println("No path found from A to E")
	}
}
