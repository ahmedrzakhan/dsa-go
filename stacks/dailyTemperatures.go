package main

import "fmt"

/**
739. Daily Temperatures

Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/

/**
brute force:
check if next item is grater than current item using 2 for loops
// TC - O(N), SC - O(1)
*/

// TC O(N), SC O(N)
func dailyTemperatures(A []int) []int {
	N := len(A)
	res := make([]int, N) // Array to store the result
	var stack []int       // Use a slice as a stack

	// Iterate through the temperatures in reverse order
	for i := N - 1; i >= 0; i-- {
		// Pop from the stack until we find a warmer temperature or the stack is empty
		for len(stack) > 0 && A[i] >= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1] // Pop from the stack
		}

		// If the stack is not empty, calculate the days until the next warmer temperature
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}

		// Push the current index onto the stack
		stack = append(stack, i)
	}

	return res
}

func mainDT() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(arr))
}
