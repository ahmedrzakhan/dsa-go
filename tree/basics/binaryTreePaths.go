package main

import "fmt"

/**
257. Binary Tree Paths

Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

Example 1:

Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
Example 2:

Input: root = [1]
Output: ["1"]

Constraints:

The number of nodes in the tree is in the range [1, 100].
-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	findPaths(root, "", &paths)
	return paths
}

// TC - O(N), SC - O(N)
func findPaths(curr *TreeNode, currPath string, paths *[]string) {
	if curr == nil {
		return
	}

	currPath += fmt.Sprintf("%d->", curr.Val)

	if curr.Left == nil && curr.Right == nil { // Leaf node
		*paths = append(*paths, currPath[:len(currPath)-2]) // Remove trailing "->"
	} else {
		findPaths(curr.Left, currPath, paths)
		findPaths(curr.Right, currPath, paths)
	}
}

func mainBTP() {
	// Example binary tree
	// root := &TreeNode{
	// 	Val:   1,
	// 	Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
	// 	Right: &TreeNode{Val: 3},
	// }

	// allPaths := binaryTreePaths(root)
	// fmt.Println(allPaths) // Output: ["1->2->5", "1->3"]

	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
		Right: &TreeNode{Val: 3},
	}

	allPaths := binaryTreePaths(root)
	fmt.Println(allPaths) // Output: [1->2->5->6 1->2->5->7 1->3]

}

/*
1
/ \
2   3
/ \
5
/ \
6   7
*/
