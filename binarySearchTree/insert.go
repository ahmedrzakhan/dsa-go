package main

import "fmt"

/**
701. Insert into a Binary Search Tree

You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

Example 1:

Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:

Example 2:

Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]
Example 3:

Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]

Constraints:

The number of nodes in the tree will be in the range [0, 104].
-108 <= Node.val <= 108
All the values Node.val are unique.
-108 <= val <= 108
It's guaranteed that val does not exist in the original BST.
*/

// TreeNode struct representing a node in the BST
type TreeNodec struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// Helper function to create a new TreeNode
func newTreeNodec(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func mainc() {
	// Example use
	root := newTreeNode(4)
	root.Left = newTreeNode(2)
	root.Right = newTreeNode(7)
	root.Left.Left = newTreeNode(1)
	root.Left.Right = newTreeNode(3)

	// Value to be inserted
	valToInsert := 5
	root = insertIntoBST(root, valToInsert)

	// Output the root value for demonstration purposes
	fmt.Println("Root value after insertion:", root.Val)
}
