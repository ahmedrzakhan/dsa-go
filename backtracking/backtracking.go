package main

import "fmt"

/**
Determine if a path exists from the root of the tree to a leaf node
it may not contain zeroes
*/

// Definition of TreeNode
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func CanReachLeaf(root *TreeNode) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if CanReachLeaf(root.Left) {
		return true
	}
	if CanReachLeaf(root.Right) {
		return true
	}
	return false
}

func LeafPath(root *TreeNode, path *[]int) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	*path = append(*path, root.Val)

	if root.Left == nil && root.Right == nil {
		return true
	}
	if LeafPath(root.Left, path) {
		return true
	}
	if LeafPath(root.Right, path) {
		return true
	}

	// remove the last element in path if both recursive calls return false
	*path = (*path)[:len(*path)-1]
	return false
}

func main() {
	// Create a binary tree
	// Example:
	//        4
	//       / \
	//      0    1
	//       \   / \
	//        7  3   2
	//            \
	//             0
	root := NewTreeNode(4)
	root.Left = NewTreeNode(0)
	root.Right = NewTreeNode(1)
	root.Left.Right = NewTreeNode(7)
	root.Right.Left = NewTreeNode(3)
	root.Right.Right = NewTreeNode(2)
	root.Right.Left.Right = NewTreeNode(0)

	// Check if there's a path to a leaf
	if CanReachLeaf(root) {
		fmt.Println("There is a path to a leaf")
	} else {
		fmt.Println("No path to a leaf")
	}

	// Find a path to a leaf
	var path []int
	if LeafPath(root, &path) {
		fmt.Println("Path to a leaf found:", path)
	} else {
		fmt.Println("No path to a leaf found")
	}
}
