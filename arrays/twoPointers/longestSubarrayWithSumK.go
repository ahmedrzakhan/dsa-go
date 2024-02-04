package main

import "fmt"

/**
https://www.geeksforgeeks.org/problems/longest-sub-array-with-sum-k0809/1

Longest Sub-Array with Sum K

banner
Given an array containing N integers and an integer K., Your task is to find the length of the longest Sub-Array with the sum of the elements equal to the given value K.

Example 1:

Input :
A[] = {10, 5, 2, 7, 1, 9}
K = 15
Output : 4
Explanation:
The sub-array is {5, 2, 7, 1}.
Example 2:

Input :
A[] = {-1, 2, 3}
K = 6
Output : 0
Explanation:
There is no such sub-array with sum 6.
Your Task:
This is a function problem. The input is already taken care of by the driver code. You only need to complete the function lenOfLongSubarr() that takes an array (A), sizeOfArray (n),  sum (K)and returns the required length of the longest Sub-Array. The driver code takes care of the printing.

Expected Time Complexity: O(N).
Expected Auxiliary Space: O(N).
*/

// TC - O(N), SC - O(N)
// Function to find the length of the longest sub-array with sum k
func lenOfLongSubarr(arr []int, N int, K int) int {
	// Map to store the earliest sum occurrences
	sumMap := make(map[int]int)
	cumSum := 0 // Initialize sum of elements
	maxLen := 0 // Initialize result

	sumMap[0] = -1
	// Traverse through the given array
	for i := 0; i < N; i++ {
		// Add the current element to the sum
		cumSum += arr[i]

		// When sub-array starts from index '0'
		// OR
		// insert 0 as value -1 in hashmap
		if cumSum == K {
			maxLen = i + 1
		}

		// Check if 'cumSum-K' is present in sumMap or not
		if prevI, exists := sumMap[cumSum-K]; exists {
			// Update maxLength
			maxLen = max(maxLen, i-prevI)
		}

		// Add cumSum to sumMap, only if it is not already present
		// This is to ensure we consider the first occurrence for a particular sum
		if _, exists := sumMap[cumSum]; !exists {
			sumMap[cumSum] = i
		}
	}

	return maxLen
}

// TC - O(N), SC - O(1)
// NOTE: works only for positive integers
func lenOfLongSubarrPos(arr []int, target int) int {
	N, sum, left, maxLength := len(arr), 0, 0, 0

	for right := 0; right < N; right++ {
		sum += arr[right]
		for sum > target {
			sum -= arr[left]
			left++
		}

		if sum == target {
			maxLength = max(maxLength, right-left+1)
		}
	}
	return maxLength
}

func maxLS(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mainLS() {
	arr := []int{10, 5, 2, 7, 1, 9}
	target := 15
	arr = []int{1, 2, -5, 1, 2}
	target = 0
	fmt.Println(lenOfLongSubarr(arr, len(arr), target))
}
