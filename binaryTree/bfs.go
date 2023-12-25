package main

import "fmt"

type TreeNodebfs struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root} // Initialize a queue with the root node

	for len(queue) > 0 {
		currentNode := queue[0] // Get the first element of the queue
		queue = queue[1:]       // Remove the first element from the queue

		// Add the current node's value to the result slice
		result = append(result, currentNode.Val)

		// Add the left and right children of the current node to the queue
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	return result
}

func mainbfs() {
	// Create a sample tree:
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}

	result := bfs(root)
	fmt.Println(result) // Output should be [1, 2, 3, 4, 5]
}
