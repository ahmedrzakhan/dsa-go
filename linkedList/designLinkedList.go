package main

import (
	"fmt"
)

/**
707. Design Linked List

Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement the MyLinkedList class:

MyLinkedList() Initializes the MyLinkedList object.
int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
void addAtTail(int val) Append a node of value val as the last element of the linked list.
void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.

Example 1:

Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3

Constraints:

0 <= index, val <= 1000
Please do not use the built-in LinkedList library.
At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex.
*/

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type MyDoublyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func Constructor() MyDoublyLinkedList {
	return MyDoublyLinkedList{}
}

func (dll *MyDoublyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
}

func (dll *MyDoublyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val}
	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
}

func (dll *MyDoublyLinkedList) Get(index int) int {
	current := dll.Head
	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}
	if current == nil {
		return -1
	}
	return current.Val
}

func (dll *MyDoublyLinkedList) DeleteAtHead() {
	if dll.Head != nil {
		if dll.Head.Next == nil {
			dll.Head = nil
			dll.Tail = nil
		} else {
			dll.Head = dll.Head.Next
			dll.Head.Prev = nil
		}
	}
}

func (dll *MyDoublyLinkedList) DeleteAtTail() {
	if dll.Tail != nil {
		if dll.Tail.Prev == nil {
			dll.Head = nil
			dll.Tail = nil
		} else {
			dll.Tail = dll.Tail.Prev
			dll.Tail.Next = nil
		}
	}
}

func (dll *MyDoublyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	if index == 0 {
		dll.AddAtHead(val)
		return
	}
	newNode := &ListNode{Val: val}
	current := dll.Head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		return
	}
	if current.Next == nil {
		dll.AddAtTail(val)
		return
	}
	newNode.Prev = current
	newNode.Next = current.Next
	current.Next.Prev = newNode
	current.Next = newNode
}

func (dll *MyDoublyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || dll.Head == nil {
		return
	}
	if index == 0 {
		dll.DeleteAtHead()
		return
	}
	current := dll.Head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		return
	}
	if current.Next == nil {
		dll.DeleteAtTail()
		return
	}
	current.Prev.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
}

func mainDll() {
	dll := Constructor()
	dll.AddAtHead(1)

	dll.AddAtTail(3)
	dll.AddAtHead(2) // List becomes 2 -> 1 -> 3
	dll.AddAtHead(4)
	fmt.Println(dll.Get(1)) // Returns 1
	// dll.DeleteAtHead()      // List becomes 1 -> 3
	// dll.DeleteAtTail()      // List becomes 1
	fmt.Println(dll.Get(0)) // Returns 1

	dll.AddAtIndex(2, 4)    // List becomes 1 -> 4
	fmt.Println(dll.Get(1)) // Returns 4
	dll.DeleteAtIndex(0)    // List becomes 4
	fmt.Println(dll.Get(0)) // Returns 4
}
