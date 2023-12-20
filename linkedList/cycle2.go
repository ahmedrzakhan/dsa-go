package main

import "fmt"

/**
142. Linked List Cycle II

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
Example 2:

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
Example 3:

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?

*/

// ListNode represents a node in a singly-linked list.
type ListNodecc struct {
	Val  int
	Next *ListNode
}

// TC - O(N), SC - O(1)
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	found := false

	// First step: use fast and slow pointers to find meeting point
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			found = true
			break
		}
	}

	// No cycle found
	if !found {
		return nil
	}

	// Second step: find the start of the cycle
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

// createCycleList creates a linked list with a cycle for testing.
func createCycleList(vals []int, cycleIndex int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	var cycleNode, tail *ListNode

	for i, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
		if i == cycleIndex {
			cycleNode = current
		}
		if i == len(vals)-1 {
			tail = current
		}
	}

	if cycleNode != nil {
		tail.Next = cycleNode // Create a cycle
	}

	return dummy.Next
}

func mainCyc() {
	// Example: Create a list with a cycle
	// List: 1 -> 2 -> 3 -> 4 -> 2 (cycle starts at index 1)
	head := createCycleList([]int{1, 2, 3, 4}, 1)

	if node := detectCycle(head); node != nil {
		fmt.Println("Cycle detected starting at node with value:", node.Val)
	} else {
		fmt.Println("No cycle detected")
	}
}
