package main

import (
	"fmt"
)

/**
230. Kth Smallest Element in a BST

Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:

Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

Constraints:

The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104

Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
*/

// TreeNode represents a node in the binary search tree.
type TreeNodeKth struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC  - O(N), SC - O(H)
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	return inOrderTraversal(root, &count, k)
}

// inOrderTraversal performs the in-order traversal of the BST.
func inOrderTraversal(node *TreeNode, count *int, k int) int {
	if node == nil {
		return -1
	}

	if node.Left != nil {
		val := inOrderTraversal(node.Left, count, k)
		if *count == k {
			return val
		}
	}

	*count++
	if *count == k {
		return node.Val
	}

	if node.Right != nil {
		return inOrderTraversal(node.Right, count, k)
	}

	return -1
}

// TC - (H+k), SC - O(H)
func kthSmallestIt(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root
	count := 0

	for node != nil || len(stack) > 0 {
		// Push all left children onto the stack
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// Process the node
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		// If count is k, return the value
		if count == k {
			return node.Val
		}

		// Move to the right child
		node = node.Right
	}

	// If k is out of bounds
	return -1
}

// helper function to create a new TreeNode.
func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// main function to test kthSmallest.
func mainKth() {
	// Example BST:
	//       3
	//      / \
	//     1   4
	//      \
	//       2
	root := newNode(3)
	root.Left = newNode(1)
	root.Right = newNode(4)
	root.Left.Right = newNode(2)

	k := 1
	// fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallestIt(root, k))
	/**
	    5
	   / \
	  3   6
	 / \
	2   4
	 /
	1
	*/

	root1 := newNode(5)
	root1.Left = newNode(3)
	root1.Right = newNode(6)
	root1.Left.Left = newNode(2)
	root1.Left.Right = newNode(4)
	root1.Left.Left.Left = newNode(1)
	// fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallestIt(root, k))

	/**
	    5
	   / \
	  3   6
	 / \
	2   4
	 /
	1
	*/
	root2 := newNode(8)
	root2.Left = newNode(3)
	root2.Right = newNode(10)
	root2.Left.Left = newNode(1)
	root2.Left.Right = newNode(6)
	root2.Right.Right = newNode(14)
	root2.Left.Right.Left = newNode(4)
	root2.Left.Right.Right = newNode(7)
	root2.Right.Right.Left = newNode(13)
	fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallestIt(root2, 3))

	/**
	    20
	   /
	  10
	 /  \
	5    15
	   /    \
	  12    18
	*/
	root3 := newNode(20)
	root3.Left = newNode(10)
	root3.Left.Left = newNode(5)
	root3.Left.Right = newNode(15)
	root3.Left.Right.Left = newNode(12)
	root3.Left.Right.Right = newNode(18)
	fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallest(root3, 3))

	/**
	2
	*/
	root4 := newNode(2)

	fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallest(root4, 3))

	/**
	1
	 \
	  2
	   \
	    3
	     \
	      4

	*/
	root5 := newNode(1)
	root5.Right = newNode(2)
	root5.Right.Right = newNode(3)
	root5.Right.Right.Right = newNode(4)
	fmt.Printf("The %dth smallest element in the BST is: %d\n", k, kthSmallest(root5, 3))

}
