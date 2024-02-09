package main

import "fmt"

/**
145. Binary Tree Postorder Traversal

Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]
Output: [3,2,1]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]


Constraints:

The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// TC - O(H), SC - O(H)
func postorderTraversal(root *TreeNode) []int {
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

func mainPO() {
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

	postorder := postorderTraversal(root) // L R P

	fmt.Println("Postorder Traversal: ", postorder) // Should print: [4 5 2 3 1]
}
