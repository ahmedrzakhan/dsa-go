package main

import "fmt"

type Node struct {
	Key int
	L   *Node
	R   *Node
}

type BST struct {
	Root *Node
}

func (t *BST) Search(key int) bool {
	curr := t.Root
	for curr != nil {
		if key == curr.Key {
			return true
		} else if key < curr.Key {
			curr = curr.L
		} else {
			curr = curr.R
		}
	}
	return false
}

func (t *BST) Delete(key int) {
	t.Root = deleteNodeBst(t.Root, key)
}

func deleteNodeBst(curr *Node, key int) *Node {
	if curr == nil {
		return nil // Key not found
	}

	if key < curr.Key {
		curr.L = deleteNodeBst(curr.L, key)
	} else if key > curr.Key {
		curr.R = deleteNodeBst(curr.R, key)
	} else { // Key found

		// Case 1: No children (leaf node)
		if curr.L == nil && curr.R == nil {
			return nil

			// Case 2: One child
		} else if curr.L == nil {
			return curr.R
		} else if curr.R == nil {
			return curr.L

			// Case 3: Two children
		} else {
			successor := findMinBst(curr.R)
			curr.Key = successor.Key
			curr.R = deleteNodeBst(curr.R, successor.Key)
		}
	}
	return curr
}

func (t *BST) Insert(key int) {
	newNode := &Node{Key: key}
	if t.Root == nil {
		t.Root = newNode
		return
	}

	curr := t.Root
	for {
		if key < curr.Key {
			if curr.L == nil {
				curr.L = newNode
				return
			}
			curr = curr.L
		} else {
			if curr.R == nil {
				curr.R = newNode
				return
			}
			curr = curr.R
		}
	}
}

// Helper function to find the inorder successor
func findMinBst(curr *Node) *Node {
	for curr.L != nil {
		curr = curr.L
	}
	return curr
}

func mainBst() {
	tree := BST{} // Create a new BST

	// Insert some values
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(6)

	// Search for a value
	fmt.Println("Is 6 in the tree?", tree.Search(6))

	// Delete a value
	tree.Delete(3)
	fmt.Println("Is 3 in the tree after deletion?", tree.Search(3))
}
