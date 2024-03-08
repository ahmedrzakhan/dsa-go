package main

import (
	"fmt"
	"math"
)

/**
530. Minimum Absolute Difference in BST

Given the root of a Binary Search Tree (BST), return the minimum absolute
difference between the values of any two different nodes in the tree.

Example 1:

Input: root = [4,2,6,1,3]
Output: 1
Example 2:

Input: root = [1,0,48,null,null,12,49]
Output: 1

Constraints:

The number of nodes in the tree is in the range [2, 104].
0 <= Node.val <= 105
*/

type TreeNodeMA struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(H)
func getMinimumDifference(curr *TreeNode) int {
	minDiff := math.MaxInt32
	prev := (*TreeNode)(nil) // Initially, no previous node

	// In-order traversal (iterative)
	stack := []*TreeNode{}

	for curr != nil || len(stack) > 0 {
		// Push left nodes onto stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Process the top node
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil {
			diff := abs(curr.Val - prev.Val)
			minDiff = min(minDiff, diff)
		}
		prev = curr

		// Move to the right
		curr = curr.Right
	}

	return minDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// (Example usage in main function)

func mainMA() {
	// Sample BST 1 (From Example 1)
	root1 := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 6}}
	/**
	      4
	     / \
	    2   6
	   / \
	  1   3
	*/
	// Sample BST 2 (From Example 2)
	root2 := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49}}}
	/**
	  1
	   \
	    48
	   / \
	  12  49
	*/
	// Test Cases
	minDiff1 := getMinimumDifference(root1)
	minDiff2 := getMinimumDifference(root2)

	fmt.Println("Test Case 1 - Minimum Difference:", minDiff1)
	fmt.Println("Test Case 2 - Minimum Difference:", minDiff2)

}
