package main

import "fmt"

/**
876. Middle of the Linked List

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Example 1:

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/

// ListNode is a node in a singly linked list.
type ListNodess struct {
	Val  int
	Next *ListNode
}

// middleNode returns the middle node of the linked list.
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// createList creates a linked list from a slice of values and returns the head of the list.
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func mainLi() {
	// Create a linked list with elements.
	vals := []int{1, 2, 3, 4, 5, 6}
	head := createList(vals)

	// Find and print the middle node value.
	midNode := middleNode(head)
	fmt.Println("The middle node value is:", midNode.Val)
}
