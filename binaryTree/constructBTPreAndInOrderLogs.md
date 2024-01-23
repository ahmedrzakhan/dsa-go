        3
       / \
      9  20
     /|  / \
    1 4 15  7
     /     / \
    2     5   8

preorder [3 9 1 4 2 20 15 7 5 8]
inorder [1 9 2 4 3 15 20 5 7 8]
rootVal 3
midIndex 4
left build tree
preorder [9 1 4 2]
inorder [1 9 2 4]
rootVal 9
midIndex 1
left build tree
preorder [1]
inorder [1]
rootVal 1
midIndex 0
left build tree
nil <nil>
root.Left <nil>
right build tree
nil <nil>
root.Right <nil>
root.Left &{1 <nil> <nil>}
right build tree
preorder [4 2]
inorder [2 4]
rootVal 4
midIndex 1
left build tree
preorder [2]
inorder [2]
rootVal 2
midIndex 0
left build tree
nil <nil>
root.Left <nil>
right build tree
nil <nil>
root.Right <nil>
root.Left &{2 <nil> <nil>}
right build tree
nil <nil>
root.Right <nil>
root.Right &{4 0x14000126180 <nil>}
root.Left &{9 0x140001260d8 0x14000126138}
right build tree
preorder [20 15 7 5 8]
inorder [15 20 5 7 8]
rootVal 20
midIndex 1
left build tree
preorder [15]
inorder [15]
rootVal 15
midIndex 0
left build tree
nil <nil>
root.Left <nil>
right build tree
nil <nil>
root.Right <nil>
root.Left &{15 <nil> <nil>}
right build tree
preorder [7 5 8]
inorder [5 7 8]
rootVal 7
midIndex 1
left build tree
preorder [5]
inorder [5]
rootVal 5
midIndex 0
left build tree
nil <nil>
root.Left <nil>
right build tree
nil <nil>
root.Right <nil>
root.Left &{5 <nil> <nil>}
right build tree
preorder [8]
inorder [8]
rootVal 8
midIndex 0
left build tree
nil <nil>
root.Left <nil>
right build tree
nil <nil>
root.Right <nil>
root.Right &{8 <nil> <nil>}
root.Right &{7 0x14000126300 0x14000126360}
root.Right &{20 0x14000126258 0x140001262b8}
Root of the constructed tree: 3


package main

import "fmt"

// TreeNode represents a node of the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree constructs a binary tree from preorder and inorder traversal data.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		fmt.Println("nil", nil)
		return nil
	}

	fmt.Println("preorder", preorder)
	fmt.Println("inorder", inorder)

	// The first element in preorder is always the root.
	rootVal := preorder[0]
	fmt.Println("rootVal", rootVal)
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder to split into left and right subtrees.
	midIndex := findIndex(inorder, rootVal)
	fmt.Println("midIndex", midIndex)

	// Recursively build the left and right subtrees.
	fmt.Println("left build tree")
	root.Left = buildTree(preorder[1:midIndex+1], inorder[:midIndex])
	fmt.Println("root.Left", root.Left)
	fmt.Println("right build tree")
	root.Right = buildTree(preorder[midIndex+1:], inorder[midIndex+1:])
	fmt.Println("root.Right", root.Right)
	return root
}

// find returns the index of the given value in the slice.
func findIndex(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func main() {
	// preorder := []int{3, 9, 20, 15, 7}
	// inorder := []int{9, 3, 15, 20, 7}

	preorder := []int{3, 9, 1, 4, 2, 20, 15, 7, 5, 8}
	inorder := []int{1, 9, 2, 4, 3, 15, 20, 5, 7, 8}
	root := buildTree(preorder, inorder)
	fmt.Println("Root of the constructed tree:", root.Val)
	// Add more print statements if you want to check the structure of the tree
}
