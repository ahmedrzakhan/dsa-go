package main

import (
	"fmt"
)

/**
111. Minimum Depth of Binary Tree

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 2
Example 2:

Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5

Constraints:

The number of nodes in the tree is in the range [0, 105].
-1000 <= Node.val <= 1000
*/

type TreeNodeMD struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(H)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Handle nodes with only one child
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// Recurse on both subtrees and return the smaller depth
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func mainMD() {
	// Example binary tree
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	minDepth := minDepth(root)
	fmt.Println("Minimum depth:", minDepth) // Output: Minimum depth: 2
}

/*
     3
    / \
   9   20
      / \
     15  7
*/
