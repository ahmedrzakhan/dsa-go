package main

import (
	"fmt"
	"math"
)

/**
300. Longest Increasing Subsequence

Given an integer array nums, return the length of the longest strictly increasing
subsequence.

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104
*/

// TC - O(2^N), SC - O(N)
func bruteForceLIS(A []int, i, prev int) int {
	if i == len(A) {
		return 0
	}

	taken := 0
	if A[i] > prev {
		taken = 1 + bruteForceLIS(A, i+1, A[i])
	}

	notTaken := bruteForceLIS(A, i+1, prev)

	return max(taken, notTaken)
}

// TC - O(N^2), SC - O(N)
func lengthOfLIS(A []int) int {
	N := len(A)
	DP := make([]int, N)

	for i := range DP {
		DP[i] = 1
	}

	for i := 1; i < N; i++ {
		for j := 0; j < i; j++ {
			if A[j] < A[i] {
				DP[i] = max(DP[i], DP[j]+1)
			}
		}
	}

	LIS := 0

	for _, val := range DP {
		LIS = max(LIS, val)
	}

	return LIS
}

func mainLIS() {
	nums := []int{10, 2, 22, 9, 33, 21, 50, 41, 60, 80}
	length := lengthOfLIS(nums)
	length = bruteForceLIS(nums, 0, math.MinInt64)
	fmt.Println("Length of LIS:", length)
}
