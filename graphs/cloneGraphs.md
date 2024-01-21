node1.Neighbors = []*Node{node2, node4} 4
node2.Neighbors = []*Node{node1, node3} //
node3.Neighbors = []*Node{node2, node4} //
node4.Neighbors = []*Node{node1, node3}


neighbor 2
neighbor 1
visited
neighbor 3
neighbor 2
visited
neighbor 4
neighbor 1
visited
neighbor 3
visited
neighbor 4
visited
Cloned Node Value: 1
Neighbor: 2
Neighbor: 4


package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

// TC - O(N), SC - O(N)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if _, found := visited[node]; found {
		fmt.Println("visited")
		return visited[node]
	}
	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, neighbor := range node.Neighbors {
		fmt.Println("neighbor", neighbor.Val)
		newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor, visited))
	}
	return newNode
}

func main() {
	// Example: create a simple graph.
	// Node 1 connects to 2 and 4, Node 2 connects to 1 and 3, Node 3 connects to 2 and 4, Node 4 connects to 1 and 3.
	node1 := &Node{Val: 1, Neighbors: []*Node{}}
	node2 := &Node{Val: 2, Neighbors: []*Node{}}
	node3 := &Node{Val: 3, Neighbors: []*Node{}}
	node4 := &Node{Val: 4, Neighbors: []*Node{}}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	// Clone the graph
	clonedGraph := cloneGraph(node1)

	fmt.Println("Cloned Node Value:", clonedGraph.Val)
	for _, neighbor := range clonedGraph.Neighbors {
		fmt.Printf("Neighbor: %v\n", neighbor.Val)
	}
}
