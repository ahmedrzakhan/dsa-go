package main

import (
	"container/heap"
	"fmt"
)

/**
703. Kth Largest Element in a Stream

Easy
Topics
Companies
Design a class to find the kth largest element in a stream. Note that it is the kth
largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream
of integers nums.
int add(int val) Appends the integer val to the stream and returns the element representing
the kth largest element in the stream.

Example 1:

Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8

Constraints:

1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
At most 104 calls will be made to add.
It is guaranteed that there will be at least k elements in the array when you
search for the kth element.
*/

type IntHeap []int // Define an IntHeap type that is a slice of ints

func (h IntHeap) Len() int {
	return len(h)
}

// Min-Heap, so Less means 'smaller than'
func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k    int
	heap *IntHeap
}

// TC - O(Log(K)), SC - O(K)
func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		// If the heap size is greater than k, remove the smallest element
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{k: k, heap: h}
}

// TC - O(NLog(K)), SC - O(K)
func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.heap, val)
	if kl.heap.Len() > kl.k {
		heap.Pop(kl.heap) // Pop the smallest element to maintain the size of the heap as k
	}
	return (*kl.heap)[0] // The root of the min-heap is the kth largest element
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func mainKL() {
	k := 3
	nums := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, nums)

	fmt.Printf("Add 3; kth largest = %d\n", kthLargest.Add(3))   // Returns 4
	fmt.Printf("Add 5; kth largest = %d\n", kthLargest.Add(5))   // Returns 5
	fmt.Printf("Add 10; kth largest = %d\n", kthLargest.Add(10)) // Returns 5
	fmt.Printf("Add 9; kth largest = %d\n", kthLargest.Add(9))   // Returns 8
	fmt.Printf("Add 4; kth largest = %d\n", kthLargest.Add(4))   // Returns 8
}
