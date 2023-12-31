package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		cur = cur.Next
	}
}

func reverseList(curr, prev *ListNode) *ListNode {
	if curr == nil {
		return prev
	}
	temp := curr.Next
	curr.Next = prev
	prev = curr
	curr = temp
	return reverseList(curr, prev)
}

// Example to run
func mainRll() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Original List:")
	printList(head)

	reversedHead := reverseList(head, nil)

	fmt.Println("Reversed List:")
	printList(reversedHead)
}
