package main

import "fmt"

/**
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest
path from the root node down to the farthest leaf node.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2

Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
*/

type TreeNodeMxD struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(H)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func mainMxD() {
	// Example binary tree
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	maxDepth := maxDepth(root)
	fmt.Println("Maximum depth:", maxDepth) // Output: Maximum depth: 3
}
