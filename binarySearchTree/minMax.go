package main

import "fmt"

// BST Node Structure
type TreeNodeM struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Find Minimum (Iterative)
func findMinM(curr *TreeNode) *TreeNode {
	if curr == nil {
		return nil
	}
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

// Find Maximum (Iterative)
func findMax(curr *TreeNode) *TreeNode {
	if curr == nil {
		return nil
	}
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr
}

// Find Minimum (Recursive)
func findMinRec(curr *TreeNode) *TreeNode {
	if curr == nil {
		return nil // Empty tree
	} else if curr.Left == nil {
		return curr // Found the leftmost node
	} else {
		return findMinRec(curr.Left) // Recurse to the left
	}
}

// Find Maximum (Recursive)
func findMaxRec(curr *TreeNode) *TreeNode {
	if curr == nil {
		return nil // Empty tree
	} else if curr.Right == nil {
		return curr // Found the rightmost node
	} else {
		return findMaxRec(curr.Right) // Recurse to the right
	}
}

func main() {
	// Sample BST (Feel free to modify)
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 8}},
		Right: &TreeNode{Val: 18}}

	minNode := findMin(root)
	maxNode := findMax(root)

	fmt.Println("Minimum Value:", minNode.Val)
	fmt.Println("Maximum Value:", maxNode.Val)

	minNode = findMinRec(root)
	maxNode = findMaxRec(root)

	fmt.Println("Minimum Value:", minNode.Val)
	fmt.Println("Maximum Value:", maxNode.Val)
}
