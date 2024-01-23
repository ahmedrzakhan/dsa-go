package main

import (
	"fmt"
)

/**
199. Binary Tree Right Side View

Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example 1:

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
Example 2:

Input: root = [1,null,3]
Output: [1,3]
Example 3:

Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/

// TreeNode struct for a binary tree node.
type TreeNodeBrs struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
// rightSideView returns the values of the nodes you can see ordered from top to bottom.
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			// If it's the last node in the current level, add it to the result.
			if i == levelSize-1 {
				result = append(result, currentNode.Val)
			}

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}

	return result
}

func mainBR() {
	// Example tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}}}
	root.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}

	fmt.Println("Right Side View:", rightSideView(root))
}

/*
1
/ \
2   3
\   \
 5   4
/
6
*/
