package main

import "fmt"

type TreeNodeID struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(H), SC - O(H)
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop

		result = append(result, curr.Val)

		curr = curr.Right
	}

	return result
}

// TC - O(H), SC - O(H)
func preorderTraversalID(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop

		result = append(result, curr.Val)

		// Push right child first so that left child is processed first
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	return result
}

// TC - O(H), SC - O(H)
func postorderTraversalID(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	stack := []*TreeNode{root}
	visit := []*TreeNode{}

	// Use the first stack to process nodes and push them to the second stack
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop from stack

		visit = append(visit, curr) // Push to visit

		// Push left and right children of removed item to stack
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	// Pop all items from visit and add to result
	for len(visit) > 0 {
		curr := visit[len(visit)-1]
		visit = visit[:len(visit)-1] // Pop from visit
		result = append(result, curr.Val)
	}

	return result
}

func mainId() {
	// Constructing a simple binary tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	inorder := inorderTraversal(root)     // L P R
	preorder := preorderTraversal(root)   // P L R
	postorder := postorderTraversal(root) // L R P

	fmt.Println("Inorder Traversal: ", inorder)     // Should print: [4 2 5 1 3]
	fmt.Println("Preorder Traversal: ", preorder)   // Should print: [1 2 4 5 3]
	fmt.Println("Postorder Traversal: ", postorder) // Should print: [4 5 2 3 1]
}
