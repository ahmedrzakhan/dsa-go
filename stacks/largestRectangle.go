package main

import "fmt"

/**
84. Largest Rectangle in Histogram

Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
return the area of the largest rectangle in the histogram.

Example 1:

Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.
Example 2:

Input: heights = [2,4]
Output: 4

Constraints:

1 <= heights.length <= 105
0 <= heights[i] <= 104
*/

// TC - O(N), SC - O(1)
func largestRectangleArea(A []int) int {
	stack := []int{} // Stack to store indices of heights
	maxArea := 0

	// Add a sentinel 0 to handle the remaining bars in the stack at the end
	A = append(A, 0)

	for idx, height := range A {
		for len(stack) > 0 && A[stack[len(stack)-1]] > height {
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop bar from the stack

			// Calculate area with the popped bar as the smallest
			width := idx
			if len(stack) > 0 {
				width = idx - stack[len(stack)-1] - 1
			}
			area := A[topIdx] * width
			maxArea = max(maxArea, area)
		}
		stack = append(stack, idx) // Push index of the current height
	}

	return maxArea
}

func mainLR() {
	heights := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(heights)
	fmt.Println("Largest Area:", area)
}
