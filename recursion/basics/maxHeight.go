package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func maxDepthF(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthF(root.Left)
	right := maxDepthF(root.Right)
	return max(left, right) + 1
	// return max(maxDepth(root.Left), maxDepth(root.Right))
}

func mainMaxHeight() {
	// Example Tree
	//        3
	//       / \
	//      9  20
	//        /  \
	//       15   7
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println("Maximum depth of the tree is:", maxDepth(root))
}
