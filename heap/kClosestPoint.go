package main

import (
	"container/heap"
	"fmt"
)

/**
973. K Closest Points to Origin

Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k,
return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).

Example 1:

Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.

Constraints:

1 <= k <= points.length <= 104
-104 <= xi, yi <= 104
*/

/**
Array Approach - TC O(N log N) due to full sorting
Heap Approach - O(N log k) due to selective heap adjustments
*/

type Point struct {
	X, Y int
}

// Implement heap.Interface for []Point (min-heap)
type PointHeap []Point

func (h *PointHeap) Len() int {
	return len(*h)
}

func (h *PointHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *PointHeap) Less(i, j int) bool {
	return (*h)[i].calculateDistance() < (*h)[j].calculateDistance()
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (point *Point) calculateDistance() float64 {
	return float64(point.X*point.X + point.Y*point.Y) // Optimized with math.Sqrt
}

// TC - O(N * LogN), SC - O(N)
func kClosest(arr [][]int, k int) [][]int {
	pq := &PointHeap{} // Create a min-heap
	heap.Init(pq)

	// Push all arr onto the heap
	for _, pointArr := range arr {
		point := Point{X: pointArr[0], Y: pointArr[1]}
		heap.Push(pq, point)
	}

	// Pop the k closest points
	var result [][]int
	for i := 0; i < k; i++ {
		point := heap.Pop(pq).(Point)
		result = append(result, []int{point.X, point.Y})
	}

	return result
}

func mainKC() {
	coordinates := [][]int{{1, 3}, {-2, 2}, {5, 1}}
	k := 2
	result := kClosest(coordinates, k)
	fmt.Println(result)
}
