package main

import (
	"container/heap"
	"fmt"
)

/**
215. Kth Largest Element in an Array

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

/**
brute force: sort and return arr len - k, TC - O(NLogN), SC - O(1)
*/

// MinHeap to store 'k' largest elements
type MinHeap []int

func (h *MinHeap) LenKL() int {
	return len(*h)
}

func (h *MinHeap) LessKL(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) SwapKL(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) PushKL(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) PopKL() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// TC  - O(NLogK), SC - O(K)
func findKthLargest(arr []int, K int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, num := range arr {
		heap.Push(minHeap, num)
		if minHeap.Len() > K {
			heap.Pop(minHeap)
		}
	}

	return heap.Pop(minHeap).(int) // kth largest at root
}

func partition(arr []int, L, R int) int {
	pivot := arr[R]
	i := L - 1

	for j := L; j < R; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[R] = arr[R], arr[i+1]
	return i + 1
}

// TC - O(N)/O(N^2), SC  - O(1)
func findKthLargestQ(arr []int, K int) int {
	N := len(arr)
	L, R := 0, N-1

	for L <= R {
		pivot := partition(arr, L, R)
		if pivot == N-K {
			return arr[pivot]
		} else if pivot > N-K {
			R = pivot - 1
		} else {
			L = pivot + 1
		}
	}
	return -1 // Not found (error handling in real problem)
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	kthLargest := findKthLargest(nums, k)
	fmt.Println(kthLargest) // Output: 4
}
