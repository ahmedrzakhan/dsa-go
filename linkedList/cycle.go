package main

import (
	"fmt"
)

/**
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
Example 2:

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/

// ListNode represents a node in a singly-linked list.
type ListNodesss struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

// createList creates a linked list and can introduce a cycle for testing.
func createListss(vals []int, cycleIndex int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	var cycleNode *ListNode

	for i, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
		if i == cycleIndex {
			cycleNode = current
		}
	}

	if cycleNode != nil {
		current.Next = cycleNode // Create a cycle
	}

	return dummy.Next
}

func mainss() {
	// Example: Create a list with a cycle
	// List: 1 -> 2 -> 3 -> 4 -> 2 (cycle starts at index 1)
	head := createListss([]int{1, 2, 3, 4}, 1)

	fmt.Println("Does the list have a cycle?", hasCycle(head))
}
