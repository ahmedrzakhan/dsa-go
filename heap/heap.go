package main

import (
	"fmt"
)

type Heap struct {
	Data []int
}

func NewHeap() *Heap {
	return &Heap{
		Data: append([]int{}, 0),
	}
}

func (h *Heap) Push(val int) {
	h.Data = append(h.Data, val)
	i := len(h.Data) - 1

	// Percolate up
	for i > 1 && h.Data[i] < h.Data[i/2] {
		// Swap
		h.Data[i], h.Data[i/2] = h.Data[i/2], h.Data[i]
		i = i / 2
	}
}

func (h *Heap) Pop() int {
	if len(h.Data) == 1 {
		return -1
	}

	if len(h.Data) == 2 {
		res := h.Data[1]
		h.Data = h.Data[:1]
		return res
	}

	res := h.Data[1]
	h.Data[1] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	i := 1

	// Percolate down
	for 2*i < len(h.Data) { // 2 < 6
		if 2*i+1 < len(h.Data) && h.Data[2*i+1] < h.Data[2*i] && h.Data[i] > h.Data[2*i+1] {
			// Swap right child
			h.Data[i], h.Data[2*i+1] = h.Data[2*i+1], h.Data[i]
			i = 2*i + 1
		} else if h.Data[i] > h.Data[2*i] {
			// Swap left child
			h.Data[i], h.Data[2*i] = h.Data[2*i], h.Data[i]
			i = 2 * i
		} else {
			break
		}
	}
	return res
}

func (h *Heap) Top() int {
	if len(h.Data) > 1 {
		return h.Data[1]
	}
	return -1
}

func (h *Heap) Heapify(arr []int) {
	// 0-th position is moved to the end
	arr = append(arr, arr[0])
	h.Data = arr
	curr := (len(h.Data) - 1) / 2

	for curr > 0 {
		// Percolate Down
		i := curr
		for 2*i < len(h.Data) {
			if 2*i+1 < len(h.Data) && h.Data[2*i+1] < h.Data[2*i] && h.Data[i] > h.Data[2*i+1] {
				// Swap right child
				h.Data[i], h.Data[2*i+1] = h.Data[2*i+1], h.Data[i]
				i = 2*i + 1
			} else if h.Data[i] > h.Data[2*i] {
				// Swap left child
				h.Data[i], h.Data[2*i] = h.Data[2*i], h.Data[i]
				i = 2 * i
			} else {
				break
			}
		}
		curr--
	}
}

func mainH() {
	heap := NewHeap()

	// Push elements into the heap
	heap.Push(10)
	heap.Push(5)
	heap.Push(15)
	heap.Push(3)

	// Print the top element
	fmt.Println("Top of the heap:", heap.Top())

	// Pop elements and print them
	for len(heap.Data) > 1 {
		fmt.Println("Popped:", heap.Pop())
		fmt.Println("Next Top:", heap.Top())
	}
}
