package main

import "fmt"

/**
input : 1 -> 2-> 3 -> 4 -> 5
output : 1 -> 2-> 5-> 4 -> 3
*/

type ListNodere struct {
	Val  int
	Next *ListNode
}

func createListre(vals []int) *ListNode {
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

func printListre(head *ListNode) *ListNode {
	curr := head

	for curr != nil {
		fmt.Print(" -> ", curr.Val)
		curr = curr.Next
	}
	return curr
}

func reverseLLre(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func getUpdatedLL(head *ListNode) *ListNode {
	curr := head
	length := 0
	for curr != nil {
		curr = curr.Next
		length++
	}

	newCurr, connectionNode := head, head

	for i := 0; i < length/2; i++ {
		connectionNode = newCurr
		newCurr = newCurr.Next
	}

	reversedHead := reverseLL(newCurr)

	connectionNode.Next = reversedHead

	return head
}

func mainrev() {
	ll := []int{1, 2, 3, 4, 5}
	head := createList(ll)
	updated := getUpdatedLL(head)
	printList(updated)
}
