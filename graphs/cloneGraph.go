package main

import (
	"fmt"
)

/**
133. Clone Graph
Medium
Topics
Companies
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}

Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the
first node with val == 1, the second node with val == 2, and so on. The graph is represented in
the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list
describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given
node as a reference to the cloned graph.

Example 1:

Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
Example 2:


Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with
val = 1 and it does not have any neighbors.
Example 3:

Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.


Constraints:

The number of nodes in the graph is in the range [0, 100].
1 <= Node.val <= 100
Node.val is unique for each node.
There are no repeated edges and no self-loops in the graph.
The Graph is connected and all nodes can be visited starting from the given node.
*/

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
	return dfsClone(node, visited)
}

func dfsClone(node *Node, visited map[*Node]*Node) *Node {
	if _, found := visited[node]; found {
		return visited[node]
	}
	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfsClone(neighbor, visited))
	}
	return newNode
}

func mainCloneGraph() {
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

	nodeWithNoNeighbors := &Node{Val: 1, Neighbors: []*Node{}}
	clonedEmptyNodeGraph := cloneGraph(nodeWithNoNeighbors)
	fmt.Println("Cloned Node Value (Empty Node):", clonedEmptyNodeGraph.Val)
	fmt.Println("Number of Neighbors (should be 0):", len(clonedEmptyNodeGraph.Neighbors))

	// Case 2: Empty graph
	emptyGraph := cloneGraph(nil)
	if emptyGraph == nil {
		fmt.Println("Cloned empty graph successfully. Result is nil, as expected.")
	} else {
		fmt.Println("Error: Expected nil for empty graph clone")
	}
}
