package main

import (
	"fmt"
	"math"
	"sort"
)

/**
628. Maximum Product of Three Numbers

Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

Example 1:

Input: nums = [1,2,3]
Output: 6
Example 2:

Input: nums = [1,2,3,4]
Output: 24
Example 3:

Input: nums = [-1,-2,-3]
Output: -6

Constraints:

3 <= nums.length <= 104
-1000 <= nums[i] <= 1000
*/

// TC - O(NLogN), SC - O(1)
func maximumProductBR(A []int) int {
	N := len(A)

	// Sort the array for easy access to min/max values
	sort.Ints(A)

	// Potential maximum candidates:
	// 1. Product of the three largest numbers
	max1 := A[N-1] * A[N-2] * A[N-3]

	// 2. Product of two smallest (negative) and the largest
	max2 := A[0] * A[1] * A[N-1]

	// Find the larger of the two candidates
	if max2 > max1 {
		return max2
	}
	return max1
}

// TC - O(N), O(1)
func maximumProduct(A []int) int {
	// Initialize values for tracking minimums and maximums

	min1, min2 := math.MaxInt64, math.MaxInt64                      // Positive infinity
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64 // Negative infinity

	for _, elem := range A {
		// Update minimum values
		if elem <= min1 {
			min2 = min1
			min1 = elem
		} else if elem <= min2 {
			min2 = elem
		}

		// Update maximum values
		if elem >= max1 {
			max3 = max2
			max2 = max1
			max1 = elem
		} else if elem >= max2 {
			max3 = max2
			max2 = elem
		} else if elem >= max3 {
			max3 = elem
		}
	}

	// Calculate potential maximum products
	maxProd1 := max1 * max2 * max3
	maxProd2 := min1 * min2 * max1

	// Return the larger maximum product
	if maxProd1 > maxProd2 {
		return maxProd1
	}
	return maxProd2
}

func mainMP() {
	nums := []int{-1, -2, -3, 4}
	maxProd := maximumProduct(nums)
	fmt.Println("Maximum Product:", maxProd)
}
