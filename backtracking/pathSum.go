package main

import "fmt"

type TreeNodePS struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
Given the root of a binary tree and an integer targetSum, return true if the tree has a
root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

Example 1:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.
Example 2:

Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.
Example 3:

Input: root = [], targetSum = 0
Output: false
Explanation: Since the tree is empty, there are no root-to-leaf paths.

Constraints:

The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
*/

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// TC - O(N), SC - O(H)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if isLeaf(root) && targetSum == root.Val {
		return true
	}
	if hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func mainPS() {
	// Example tree:
	//       5
	//      / \
	//     4   8
	//    /   / \
	//   11  13  4
	//  /  \      \
	// 7    2      1
	root := newNode(5)
	root.Left = newNode(4)
	root.Right = newNode(8)
	root.Left.Left = newNode(11)
	root.Left.Left.Left = newNode(7)
	root.Left.Left.Right = newNode(2)
	root.Right.Left = newNode(13)
	root.Right.Right = newNode(4)
	root.Right.Right.Right = newNode(1)
	fmt.Println(hasPathSum(root, 22))
}
