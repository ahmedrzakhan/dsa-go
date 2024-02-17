package main

import (
	"container/heap"
	"fmt"
)

/**
480. Sliding Window Median

The median is the middle value in an ordered integer list. If the size of the list is even,
there is no middle value. So the median is the mean of the two middle values.

For examples, if arr = [2,3,4], the median is 3.
For examples, if arr = [1,2,3,4], the median is (2 + 3) / 2 = 2.5.
You are given an integer array nums and an integer k. There is a sliding window of size k
which is moving from the very left of the array to the very right. You can only see the k
numbers in the window. Each time the sliding window moves right by one position.

Return the median array for each window in the original array. Answers within 10-5 of the
actual value will be accepted.

Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
Explanation:
Window position                Median
---------------                -----
[1  3  -1] -3  5  3  6  7        1
 1 [3  -1  -3] 5  3  6  7       -1
 1  3 [-1  -3  5] 3  6  7       -1
 1  3  -1 [-3  5  3] 6  7        3
 1  3  -1  -3 [5  3  6] 7        5
 1  3  -1  -3  5 [3  6  7]       6
Example 2:

Input: nums = [1,2,3,4,2,3,1,4,2], k = 3
Output: [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]

Constraints:

1 <= k <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
*/

// MinHeap implementation
type MinHeapSWM []int

func (h *MinHeap) LenSWM() int {
	return len(*h)
}

func (h *MinHeap) LessSWM(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) SwapSWM(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) PushSWM(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) PopSWM() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap implementation
type MaxHeapSWM []int

func (h *MaxHeap) LenSWM() int {
	return len(*h)
}

func (h *MaxHeap) LessSWM(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) SwapSWM(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) PushSWM(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) PopSWM() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Helper function to remove a number from either heap
func removeFromHeap(maxHeap *MaxHeap, minHeap *MinHeap, num int) {
	if num <= (*maxHeap)[0] {
		for i := 0; i < maxHeap.Len(); i++ {
			if (*maxHeap)[i] == num {
				heap.Remove(maxHeap, i)
				break
			}
		}
	} else {
		for i := 0; i < minHeap.Len(); i++ {
			if (*minHeap)[i] == num {
				heap.Remove(minHeap, i)
				break
			}
		}
	}

	// Rebalance heaps
	if maxHeap.Len() < minHeap.Len() {
		heap.Push(maxHeap, heap.Pop(minHeap))
	} else if maxHeap.Len() > minHeap.Len()+1 {
		heap.Push(minHeap, heap.Pop(maxHeap))
	}
}

// TC - O(NK), SC - O(K)
func medianSlidingWindow(arr []int, K int) []float64 {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)

	var medians []float64

	// Process initial window
	for i := 0; i < K; i++ {
		heap.Push(maxHeap, arr[i])
		heap.Push(minHeap, heap.Pop(maxHeap))
		if maxHeap.Len() < minHeap.Len() {
			heap.Push(maxHeap, heap.Pop(minHeap))
		}
	}
	medians = append(medians, getMedian(maxHeap, minHeap))

	// Slide the window
	for i := K; i < len(arr); i++ {
		removeFromHeap(maxHeap, minHeap, arr[i-K])

		heap.Push(maxHeap, arr[i])
		heap.Push(minHeap, heap.Pop(maxHeap))
		if maxHeap.Len() < minHeap.Len() {
			heap.Push(maxHeap, heap.Pop(minHeap))
		}

		medians = append(medians, getMedian(maxHeap, minHeap))
	}

	return medians
}

func getMedian(maxHeap *MaxHeap, minHeap *MinHeap) float64 {
	if maxHeap.Len() == minHeap.Len() {
		return float64((*maxHeap)[0]+(*minHeap)[0]) / 2.0
	}
	return float64((*maxHeap)[0])
}

func mainSWM() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	K := 3
	medians := medianSlidingWindow(nums, K)
	fmt.Println("Sliding window medians:", medians)
}
