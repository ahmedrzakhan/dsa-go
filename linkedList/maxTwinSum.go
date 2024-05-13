package main

import "fmt"

/**
2130. Maximum Twin Sum of a Linked List

In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.

Given the head of a linked list with even length, return the maximum twin sum of the linked list.

Example 1:

Input: head = [5,4,2,1]
Output: 6
Explanation:
Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
There are no other nodes with twins in the linked list.
Thus, the maximum twin sum of the linked list is 6.
Example 2:

Input: head = [4,2,2,3]
Output: 7
Explanation:
The nodes with twins present in this linked list are:
- Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
- Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
Thus, the maximum twin sum of the linked list is max(7, 4) = 7.
Example 3:

Input: head = [1,100000]
Output: 100001
Explanation:
There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.

Constraints:

The number of nodes in the list is an even integer in the range [2, 105].
1 <= Node.val <= 105
*/

// ListNode represents a node in a singly-linked list.
type ListNodedd struct {
	Val  int
	Next *ListNode
}

// TC - O(N), SC - O(1)
func pairSum(head *ListNode) int {
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	// If the list has an odd number of elements, skip the middle element
	if fast != nil {
		slow = slow.Next
	}

	// 'prev' now points to the head of the reversed first half,
	// 'slow' points to the head of the second half
	first, second := prev, slow
	maxSum := first.Val
	for first != nil && second != nil {
		maxSum = max(maxSum, first.Val+second.Val)
		first = first.Next
		second = second.Next
	}
	return maxSum
}

// createList creates a linked list from a slice of values and returns the head of the list.
func createListd(vals []int) *ListNode {
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

func maind() {
	// Example usage:
	vals := []int{4, 2, 2, 3} // Create a test list of values.
	// vals := []int{5, 4, 2, 1}
	head := createList(vals) // Create a list from those values.

	// Call pairSum and print the result.
	fmt.Println("Maximum twin sum:", pairSum(head))
}
