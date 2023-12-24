package main

import "fmt"

/**
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.

Example 1:

Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:

Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.
Example 3:

Input: root = [], key = 0
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 104].
-105 <= Node.val <= 105
Each node has a unique value.
root is a valid binary search tree.
-105 <= key <= 105

Follow up: Could you solve it with time complexity O(height of tree)?
*/

// TreeNode struct for a BST node
type TreeNodeD struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to find the minimum value in a BST
func findMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Function to delete a node from a BST
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}

	return root
}

// Helper function to create a new TreeNode
func newTreeNodeD(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func mainD() {
	// Example use
	root := newTreeNode(5)
	root.Left = newTreeNode(3)
	root.Right = newTreeNode(6)
	root.Left.Left = newTreeNode(2)
	root.Left.Right = newTreeNode(4)
	root.Right.Right = newTreeNode(7)

	key := 3
	root = deleteNode(root, key)

	// Output the root value for demonstration purposes
	fmt.Println("Root value after deletion:", root.Val)
}
