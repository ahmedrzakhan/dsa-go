package main

import (
	"fmt"
)

// Stack type based on a slice of integers
type Stacker struct {
	Data []int
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

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.Data)
}

// midDel removes the middle element of the stack
func midDel(s *Stack) {
	if s.Size() == 0 {
		return
	}

	// For an even number of elements, the middle element closer to the top is removed
	k := (s.Size() + 1) / 2
	solve(s, k)
}

// solve is a helper function for midDel
func solve(s *Stack, k int) {
	if k == 1 {
		s.Pop()
		return
	}

	temp := s.Pop()
	solve(s, k-1)
	s.Push(temp)
}

func mainMidDel() {
	st := Stack{Data: []int{1, 2, 3, 4, 5}}
	fmt.Println("Stack before deletion:", st.Data)

	midDel(&st)
	fmt.Println("Stack after deletion:", st.Data)

	st2 := Stack{Data: []int{1, 2, 3, 4, 5, 6}}
	fmt.Println("Stack before deletion:", st2.Data)

	midDel(&st2)
	fmt.Println("Stack after deletion:", st2.Data)
}
