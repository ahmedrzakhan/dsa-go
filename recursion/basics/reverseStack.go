package main

import (
	"fmt"
)

// Stack type based on a slice of integers
type StackRev struct {
	Data []int
}

// Push adds an element to the top of the stack
func (s *Stack) PushRev(v int) {
	s.Data = append(s.Data, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) PopRev() int {
	if len(s.Data) == 0 {
		panic("pop from empty stack")
	}

	element := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return element
}

// Size returns the size of the stack
func (s *Stack) SizeRev() int {
	return len(s.Data)
}

// reverse reverses the stack using recursion
func reverse(s *Stack) {
	if s.Size() == 1 {
		return
	}

	temp := s.Pop()
	reverse(s)
	insertAtBottom(s, temp)
}

// insertAtBottom inserts an element at the bottom of the stack
func insertAtBottom(s *Stack, element int) {
	if s.Size() == 0 {
		s.Push(element)
		return
	}

	temp := s.Pop()
	insertAtBottom(s, element)
	s.Push(temp)
}

func main() {
	st := Stack{Data: []int{5, 4, 3, 2, 1}}
	fmt.Println("Stack before reversing:", st.Data)

	reverse(&st)
	fmt.Println("Stack after reversing:", st.Data)
}
