package main

import (
	"container/heap"
	"fmt"
)

/**
502. IPO

Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares
to Venture Capital, LeetCode would like to work on some projects to increase its
capital before the IPO. Since it has limited resources, it can only finish at most k
distinct projects before the IPO. Help LeetCode design the best way to maximize its
total capital after finishing at most k distinct projects.

You are given n projects where the ith project has a pure profit profits[i] and a
minimum capital of capital[i] is needed to start it.

Initially, you have w capital. When you finish a project, you will obtain its pure
profit and the profit will be added to your total capital.

Pick a list of at most k distinct projects from given projects to maximize your final
capital, and return the final maximized capital.

The answer is guaranteed to fit in a 32-bit signed integer.

Example 1:

Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
Output: 4
Explanation: Since your initial capital is 0, you can only start the project indexed 0.
After finishing it you will obtain profit 1 and your capital becomes 1.
With capital 1, you can either start the project indexed 1 or the project indexed 2.
Since you can choose at most 2 projects, you need to finish the project indexed 2 to
get the maximum capital.
Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.
Example 2:

Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
Output: 6

Constraints:

1 <= k <= 105
0 <= w <= 109
n == profits.length
n == capital.length
1 <= n <= 105
0 <= profits[i] <= 104
0 <= capital[i] <= 109
*/

// Project struct to store capital and profit
type Project struct {
	capital int
	profit  int
}

// MinHeap for prioritizing projects by capital (ascending)
type MinHeapI []Project

func (h *MinHeapI) Len() int {
	return len(*h)
}

func (h *MinHeapI) Less(i, j int) bool {
	return (*h)[i].capital < (*h)[j].capital
}

func (h *MinHeapI) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeapI) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MinHeapI) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap for prioritizing projects by profit (descending)
type MaxHeapI []Project

func (h *MaxHeapI) Len() int {
	return len(*h)
}

func (h *MaxHeapI) Less(i, j int) bool {
	return (*h)[i].profit > (*h)[j].profit
}

func (h *MaxHeapI) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeapI) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MaxHeapI) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TC - O(NLogN), SC - O(N)
func findMaximizedCapital(K int, W int, profits []int, capital []int) int {
	N := len(profits)

	// Combine project data into min-heap prioritized by capital
	minCapital := &MinHeapI{}
	heap.Init(minCapital)
	for i := 0; i < N; i++ {
		heap.Push(minCapital, Project{capital[i], profits[i]})
	}

	// Max-heap to store potential projects prioritized by profit
	maxProfit := &MaxHeapI{}
	heap.Init(maxProfit)

	currentCapital := W

	// Select most profitable projects iteratively
	for i := 0; i < K; i++ {
		// Move affordable projects to maxProfit heap
		for minCapital.Len() > 0 && (*minCapital)[0].capital <= currentCapital {
			heap.Push(maxProfit, heap.Pop(minCapital).(Project))
		}

		// If no affordable projects, stop
		if maxProfit.Len() == 0 {
			break
		}

		// Select best profit opportunity
		currentCapital += heap.Pop(maxProfit).(Project).profit
	}

	return currentCapital
}

func mainI() {
	// Example usage
	k := 2
	W := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}

	result := findMaximizedCapital(k, W, profits, capital)
	fmt.Println("Maximum Capital:", result)
}
