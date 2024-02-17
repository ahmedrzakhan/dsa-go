package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

type MinHeapTH []int

func (h *MinHeap) LenTH() int {
	return len(*h)
}

func (h *MinHeap) LessTH(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) SwapTH(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) PopTH() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) PushTH(x interface{}) {
	*h = append(*h, x.(int))
}

type MaxHeap []int

func (h *MaxHeap) LenTH() int {
	return len(*h)
}

func (h *MaxHeap) LessTH(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) SwapTH(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) PopTH() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) PushTH(x interface{}) {
	*h = append(*h, x.(int))
}

func ConstructorTH() MedianFinder {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{minHeap, maxHeap}
}

func (mf *MedianFinder) AddNumTH(num int) {
	heap.Push(mf.maxHeap, num)
	val := heap.Pop(mf.maxHeap).(int)
	heap.Push(mf.minHeap, val)
	// Ensure maxHeap always has at least as many elements as minHeap
	if mf.maxHeap.Len() < mf.minHeap.Len() {
		val := heap.Pop(mf.minHeap).(int)
		heap.Push(mf.maxHeap, val)
	}
}

// TC - O(NlogN), SC - O(N)
func (mf *MedianFinder) FindMedianTH() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	} else if mf.maxHeap.Len() < mf.minHeap.Len() {
		return float64((*mf.minHeap)[0])
	}
	return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
}

func mainTH() {
	mf := ConstructorTH()

	// Example stream of numbers:
	stream := []int{5, 2, 8, 1, 10}

	for _, num := range stream {
		mf.AddNum(num)
		fmt.Printf("Number Added: %d, Current Median: %.1f\n", num, mf.FindMedian())
	}
}
