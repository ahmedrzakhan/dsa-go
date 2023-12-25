package main

import (
	"fmt"
)

/**
94. Binary Tree Inorder Traversal

Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// TreeNode is a struct for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(N), SC - O(N)
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, result)          // Visit left subtree
	*result = append(*result, node.Val) // Visit root
	inorder(node.Right, result)         // Visit right subtree
}

// TC - O(N), SC - O(N)
func inorderTraversalIt(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func mainIno() {
	// Example tree:
	//     1
	//      \
	//       2
	//      /
	//     3
	// root1 := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	// result1 := inorderTraversal(root1)
	// fmt.Println(result1)

	//     4
	//    / \
	//   2   5
	//  / \
	// 1   3
	root2 := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	result2 := inorderTraversalIt(root2)
	// result2 := inorderTraversal(root2)
	fmt.Println(result2) // Output: [1, 2, 3, 4, 5]

	// 	5
	// 	/
	//    4
	//   /
	//  3

	// root3 := &TreeNode{5, &TreeNode{4, &TreeNode{3, nil, nil}, nil}, nil}
	// result3 := inorderTraversal(root3)
	// fmt.Println(result3) // Output: [3, 4, 5]

	// 1
	// / \
	// 2   3
	// \
	//  4
	//   \
	//    5

	root4 := &TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}, &TreeNode{3, nil, nil}}
	result4 := inorderTraversal(root4)
	fmt.Println(result4) // Output: [2, 4, 5, 1, 3]

	// 1
	// \
	//  2
	// / \
	// 3   4
	//   /
	//  5

	root5 := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, &TreeNode{5, nil, nil}, nil}}}
	result5 := inorderTraversal(root5)
	fmt.Println(result5) // Output: [1, 3, 2, 5, 4]

}
