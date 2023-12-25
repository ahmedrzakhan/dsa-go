package main

import "fmt"

type TreeNodeDfs struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) []int {
	var result []int
	dfsHelper(root, &result)
	return result
}

func dfsHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	// Pre-order: process the root, then left, then right
	*result = append(*result, node.Val) // Process the root
	dfsHelper(node.Left, result)        // Traverse left subtree
	dfsHelper(node.Right, result)       // Traverse right subtree
}

func main() {
	// Example tree:
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}

	result := dfs(root)
	fmt.Println(result) // Output should be [1, 2, 4, 5, 3] for pre-order DFS
}
