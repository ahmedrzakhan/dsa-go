package main

import (
	"fmt"
)

// Stack type based on a slice of integers
type Stack struct {
	Data []int
}

// NewStack creates and returns a new Stack instance
func NewStack() Stack {
	return Stack{
		Data: []int{},
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(v int) {
	s.Data = append(s.Data, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() int {
	if len(s.Data) == 0 {
		panic("pop from empty stack")
	}

	element := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return element
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() int {
	return s.Data[len(s.Data)-1]
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.Data)
}

// sortStack recursively sorts the stack
func sortStack(s *Stack) {
	if s.Size() <= 1 {
		return
	}

	temp := s.Pop()
	sortStack(s)
	insertElementAtStack(s, temp)
}

// insertElementAtStack recursively inserts an element into the sorted stack
func insertElementAtStack(s *Stack, temp int) {
	if s.Size() == 0 || temp >= s.Peek() {
		s.Push(temp)
		return
	}

	val := s.Pop()
	insertElementAtStack(s, temp)
	s.Push(val)
}

func mainSortStack() {
	st := NewStack()
	st.Push(5)
	st.Push(1)
	st.Push(0)
	st.Push(2)

	fmt.Println("Stack before sorting:")
	for _, v := range st.Data {
		fmt.Print(v, " ")
	}

	fmt.Println("\nStack after sorting:")
	sortStack(&st)
	for _, v := range st.Data {
		fmt.Print(v, " ")
	}
}
