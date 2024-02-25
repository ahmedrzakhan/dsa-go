package main

import (
	"fmt"
	"math"
)

/**
124. Binary Tree Maximum Path Sum

A path in a binary tree is a sequence of nodes where each pair of adjacent nodes
in the sequence has an edge connecting them. A node can only appear in the sequence
at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.

Example 1:

Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
Example 2:

Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

Constraints:

The number of nodes in the tree is in the range [1, 3 * 104].
-1000 <= Node.val <= 1000
*/

// TreeNode represents a node in a binary tree
type TreeNodeMPS struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(H)
// Helper function to recursively calculate the maximum path sum
func dfs(curr *TreeNode, maxSum *int) int {
	if curr == nil {
		return 0
	}

	leftSum := max(dfs(curr.Left, maxSum), 0)   // ensuring non-negative contributions
	rightSum := max(dfs(curr.Right, maxSum), 0) // ensuring non-negative contributions

	currentSum := curr.Val + leftSum + rightSum

	*maxSum = max(*maxSum, currentSum)

	return curr.Val + max(leftSum, rightSum)
}

func maxPathSum(curr *TreeNode) int {
	maxSum := math.MinInt64 // Initialize with the minimum possible integer

	dfs(curr, &maxSum)

	return maxSum
}

func main() {
	// Example usage:
	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(maxPathSum(root)) // Output: 42
}

/**
    -10
   /  \
  9    20
      / \
     15  7
*/
