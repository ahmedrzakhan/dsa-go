package main

import "fmt"

/**
105. Construct Binary Tree from Preorder and Inorder Traversal

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary
tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]

Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
*/

// TreeNode represents a node of the binary tree.
type TreeNodeCon struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N^2), SC - O(N)
// buildTree constructs a binary tree from preorder and inorder traversal data.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// The first element in preorder is always the root.
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder to split into left and right subtrees.
	midIndex := findIndex(inorder, rootVal)

	// Recursively build the left and right subtrees.
	root.Left = buildTree(preorder[1:midIndex+1], inorder[:midIndex])
	root.Right = buildTree(preorder[midIndex+1:], inorder[midIndex+1:])
	return root
}

// find returns the index of the given value in the slice.
func findIndex(A []int, value int) int {
	for i, v := range A {
		if v == value {
			return i
		}
	}
	return -1
}

func mainCons() {
	// preorder := []int{3, 9, 20, 15, 7}
	// inorder := []int{9, 3, 15, 20, 7}

	preorder := []int{3, 9, 1, 4, 2, 20, 15, 7, 5, 8}
	inorder := []int{1, 9, 2, 4, 3, 15, 20, 5, 7, 8}
	root := buildTree(preorder, inorder)
	fmt.Println("Root of the constructed tree:", root.Val)
	// Add more print statements if you want to check the structure of the tree
}
