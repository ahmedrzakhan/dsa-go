package main

import "fmt"

/**
173. Binary Search Tree Iterator

Implement the BSTIterator class that represents an iterator over the in-order traversal of a
binary search tree (BST):

BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST
is given as part of the constructor. The pointer should be initialized to a non-existent number
smaller than any element in the BST.
boolean hasNext() Returns true if there exists a number in the traversal to the right of the
pointer, otherwise returns false.
int next() Moves the pointer to the right, then returns the number at the pointer.
Notice that by initializing the pointer to a non-existent smallest number, the first call to
next() will return the smallest element in the BST.

You may assume that next() calls will always be valid. That is, there will be at least a next
number in the in-order traversal when next() is called.

Example 1:

Input
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next",
"hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]

Explanation
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // return 3
bSTIterator.next();    // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 20
bSTIterator.hasNext(); // return False

Constraints:

The number of nodes in the tree is in the range [1, 105].
0 <= Node.val <= 106
At most 105 calls will be made to hasNext, and next.

Follow up:

Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory,
where h is the height of the tree?
*/

// TreeNode structure for a binary tree node.
type TreeNodeBST struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator structure holds the state of the iterator.
type BSTIterator struct {
	stack []*TreeNode
}

// TC - O(H), SC - O(H)

// Constructor to initialize the BSTIterator.
func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushAllLeft(root)
	return it
}

// pushAllLeft pushes all the left children of a node onto the stack.
func (bst *BSTIterator) pushAllLeft(curr *TreeNode) {
	for curr != nil {
		bst.stack = append(bst.stack, curr)
		curr = curr.Left
	}
}

// Next returns the next smallest number in the BST.
func (bst *BSTIterator) Next() int {
	// Pop the top element from the stack.
	curr := bst.stack[len(bst.stack)-1]
	bst.stack = bst.stack[:len(bst.stack)-1]
	// If there is a right child, push all its left children onto the stack.
	if curr.Right != nil {
		bst.pushAllLeft(curr.Right)
	}
	return curr.Val
}

// HasNext returns true if there is a next smallest number in the BST.
func (bst *BSTIterator) HasNext() bool {
	return len(bst.stack) > 0
}

// Main function to demonstrate the BSTIterator usage.
func mainBST() {
	// Example tree:
	//      7
	//     / \
	//    3   15
	//       /  \
	//      9   20
	root := &TreeNode{7,
		&TreeNode{3, nil, nil},
		&TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}

	it := Constructor(root)
	for it.HasNext() {
		fmt.Println(it.Next()) // Prints the elements of the BST in ascending order.
	}
}
