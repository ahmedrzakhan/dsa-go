package main

import (
	"fmt"
	"sort"
)

/**
Given the root of a binary tree, calculate the vertical order traversal of the binary tree.

For each node at position (row, col), its left and right children will be at positions (row + 1,
col - 1) and (row + 1, col + 1) respectively. The root of the tree is at (0, 0).

The vertical order traversal of a binary tree is a list of top-to-bottom orderings for each column
index starting from the leftmost column and ending on the rightmost column. There may be multiple
nodes in the same row and same column. In such a case, sort these nodes by their values.

Return the vertical order traversal of the binary tree.

Input: root = [3,9,20,null,null,15,7]
Output: [[9],[3,15],[20],[7]]

Input: root = [1,2,3,4,5,6,7]
Output: [[4],[2],[1,5,6],[3],[7]]

Input: root = [1,2,3,4,6,5,7]
Output: [[4],[2],[1,5,6],[3],[7]]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// helper function to perform DFS and record the node positions
func dfs(node *TreeNode, row, col int, nodes *[][3]int) {
	if node == nil {
		return
	}
	*nodes = append(*nodes, [3]int{col, row, node.Val})
	dfs(node.Left, row+1, col-1, nodes)
	dfs(node.Right, row+1, col+1, nodes)
}

// main function to calculate vertical order traversal
func verticalOrderTraversal(root *TreeNode) [][]int {
	nodes := make([][3]int, 0)
	dfs(root, 0, 0, &nodes)

	// Sort nodes by column, then row, then value
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i][0] != nodes[j][0] {
			return nodes[i][0] < nodes[j][0]
		}
		if nodes[i][1] != nodes[j][1] {
			return nodes[i][1] < nodes[j][1]
		}
		return nodes[i][2] < nodes[j][2]
	})

	result := [][]int{}
	lastCol := nodes[0][0]
	colValues := []int{}

	for _, node := range nodes {
		col, _, val := node[0], node[1], node[2]
		if col != lastCol {
			result = append(result, colValues)
			colValues = []int{}
			lastCol = col
		}
		colValues = append(colValues, val)
	}
	result = append(result, colValues)

	return result
}

func mainVot() {
	// Example tree:    3
	//                 / \
	//                9  20
	//                   / \
	//                  15  7
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(verticalOrderTraversal(root)) // Output: [[9],[3,15],[20],[7]]
}
