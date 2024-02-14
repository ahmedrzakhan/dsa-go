package main

import (
	"container/heap"
	"fmt"
)

/**
1046. Last Stone Weight

You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and
smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The
result of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.

Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
Example 2:

Input: stones = [1]
Output: 1

Constraints:

1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/

/**
Note: deleting an item in array is O(N) SC and = its O(LogN)
Less method > and < decide if its max heap or min heap
*/

// An IntHeap is a max-heap of ints.
type IntHeapLSW []int

func (h *IntHeap) LenLSW() int {
	return len(*h)
}

// We want Pop to give us the largest element, so we use greater than here.
func (h *IntHeap) LessLSW(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) SwapLSW(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) PushLSW(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) PopLSW() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TC  - O(NLogN), SC - O(N)
// lastStoneWeight calculates the weight of the last stone.
func lastStoneWeight(arr []int) int {
	// 1. Convert input array to an IntHeap (max-heap):
	h := IntHeap(arr)

	// 2. Initialize the heap (build the heap structure):
	heap.Init(&h)

	// 3. Iteratively compare and combine the heaviest stones:
	for h.Len() > 1 {
		// 3.1 Extract the two heaviest stones (max values):
		first := heap.Pop(&h).(int)
		second := heap.Pop(&h).(int)

		// 3.2 Calculate the weight difference:
		diff := first - second

		// 3.3 Add the remaining weight (if heavier) back to the heap:
		if diff > 0 {
			heap.Push(&h, diff)
		}
	}

	// 4. Get the final stone weight:
	if h.Len() == 0 {
		return 0 // No stones remaining
	}

	return heap.Pop(&h).(int) // Single stone remaining
}

func mainLSW() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Printf("The weight of the last stone is: %d\n", lastStoneWeight(stones))
}
