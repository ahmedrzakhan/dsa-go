package main

import "fmt"

/**
404. Sum of Left Leaves

Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
Example 2:

Input: root = [1]
Output: 0

Constraints:

The number of nodes in the tree is in the range [1, 1000].
-1000 <= Node.val <= 1000
*/

type TreeNodeSLL struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	helper(root, &sum, false) // Start at the root
	return sum
}

// TC - O(N) ,SC - O(H)
func helper(node *TreeNode, sum *int, isLeft bool) {
	if node == nil {
		return
	}

	if isLeft && node.Left == nil && node.Right == nil { // Left leaf node
		*sum += node.Val
	}

	helper(node.Left, sum, true)   // Traverse left subtree with isLeft = true
	helper(node.Right, sum, false) // Traverse right subtree with isLeft = false
}

func mainSLL() {
	// Example binary tree
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	result := sumOfLeftLeaves(root)
	fmt.Println("Sum of left leaves:", result) // Output: Sum of left leaves: 24
}
