package main

import "fmt"

/**
918. Maximum Sum Circular Subarray

Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.

A circular array means the end of the array connects to the beginning of the array. Formally, the next element of
 nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].

A subarray may only include each element of the fixed buffer nums at most once. Formally, for a subarray nums[i],
nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.

Example 1:

Input: nums = [1,-2,3,-2]
Output: 3
Explanation: Subarray [3] has maximum sum 3.
Example 2:

Input: nums = [5,-3,5]
Output: 10
Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
Example 3:

Input: nums = [-3,-2,-3]
Output: -2
Explanation: Subarray [-2] has maximum sum -2.
*/

func maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TC - O(N), SC - O(1)
func maxSubarraySumCircular(arr []int) int {
	globalMax, globalMin := arr[0], arr[0]
	currentMax, currentMin := 0, 0
	total := 0
	for _, num := range arr {
		currentMax = maxs(num, currentMax+num)
		currentMin = min(num, currentMin+num)
		total = total + num
		globalMax = maxs(globalMax, currentMax)
		globalMin = min(globalMin, currentMin)
	}

	if globalMax > 0 {
		return max(globalMax, total-globalMin)
	}
	return globalMax
}

func mainMax() {
	arr := []int{5, 4, -10, 9, -3, 2, 5, -20}
	fmt.Println(maxSubarraySumCircular(arr))
}
