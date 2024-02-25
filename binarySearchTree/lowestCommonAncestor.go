package main

import "fmt"

/**
235. Lowest Common Ancestor of a Binary Search Tree

Given a binary search tree (BST), find the lowest common ancestor (LCA) node
of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor
is defined between two nodes p and q as the lowest node in T that has both p
and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant
of itself according to the LCA definition.
Example 3:

Input: root = [2,1], p = 2, q = 1
Output: 2

Constraints:

The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the BST.
*/

// TreeNode Definition
type TreeNodeLCA struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC - O(H), SC - O(H)
// Finds the Lowest Common Ancestor
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base case: Empty tree
	if root == nil {
		return nil
	}

	// Exploit BST property:
	// * If both p and q are smaller than root, LCA is in the left subtree
	// * If both p and q are larger than root, LCA is in the right subtree
	// * Otherwise, the current root is the LCA

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

func mainLCA() {
	// Example BST
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}

	p := root.Left
	q := root.Right

	lca := lowestCommonAncestor(root, p, q)
	fmt.Println("LCA of", p.Val, "and", q.Val, "is:", lca.Val)
}

/*
       6
      / \
     2   8
    / \  / \
   0  4 7   9
*/
