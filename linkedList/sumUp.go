package main

import "fmt"

/**
asked in one of the interviews
two lists given 1-> 9 -> 8 -> 2, 2 -> 2-> 3
return sum of this in new LL
2 -> 2 -> 0 -> 5
*/

type ListNodeSum struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) Append(data int) {
	newNode := &ListNode{Val: data}

	curr := ll.Head
	if curr == nil {
		ll.Head = newNode
		return
	}

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func createListSum(vals []int) *ListNode {
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

func reverseLL(head *ListNode) *ListNode {
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

func printLL(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(" -> ", curr.Val)
		curr = curr.Next
	}
}

func getSummedList(firstHead *ListNode, secondHead *ListNode) *ListNode {
	firstReversed := reverseLL(firstHead)
	secondReversed := reverseLL(secondHead)
	// printLL(firstReversed)
	// printLL(secondReversed)

	currFirst := firstReversed
	currSecond := secondReversed

	newList := LinkedList{}

	carry := 0

	for currFirst != nil && currSecond != nil {
		sum := currFirst.Val + currSecond.Val
		if carry != 0 {
			sum = sum + 1
			carry = 0
		}
		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		}
		newList.Append(sum)
		currFirst = currFirst.Next
		currSecond = currSecond.Next
	}

	if currFirst != nil {
		for currFirst != nil {
			sum := currFirst.Val + carry
			carry = 0
			newList.Append(sum)
			currFirst = currFirst.Next
		}
	} else {
		for currSecond != nil {
			sum := currSecond.Val + carry
			carry = 0
			newList.Append(sum)
			currSecond = currSecond.Next
		}
	}

	reversedNewList := reverseLL(newList.Head)

	return reversedNewList
}

func mainsum() {
	firstLL := []int{1, 9, 8, 2}
	secondLL := []int{2, 2, 3}
	firstHead := createList(firstLL)
	secondHead := createList(secondLL)
	reversedHead := getSummedList(firstHead, secondHead)
	printLL(reversedHead)
}
