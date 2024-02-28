package main

import (
	"fmt"
	"math"
)

/**
98. Validate Binary Search Tree

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
 of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:


Input: root = [2,1,3]
Output: true
Example 2:


Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
*/

type TreeNodeVbst struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(curr *TreeNode, min int, max int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= min || curr.Val >= max {
		return false
	}

	return isValid(curr.Left, min, curr.Val) && isValid(curr.Right, curr.Val, max)
}

func mainVbst() {
	root := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 8}}
	result := isValidBST(root)
	fmt.Println(result) // Should print "true"

	root = &TreeNode{Val: 24,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 33}}},
		Right: &TreeNode{Val: 26}}
	result = isValidBST(root)
	fmt.Println(result) // Should print "true"
}

/**
	5
  /   \
  2    8


   24
  /  \
 2   26
/ \
1  5
	\
	33
*/
