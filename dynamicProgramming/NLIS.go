package main

import "fmt"

/**
673. Number of Longest Increasing Subsequence

Given an integer array nums, return the number of longest increasing subsequences.

Notice that the sequence has to be strictly increasing.

Example 1:

Input: nums = [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
Example 2:

Input: nums = [2,2,2,2,2]
Output: 5
Explanation: The length of the longest increasing subsequence is 1, and there are 5
increasing subsequences of length 1, so output 5.

Constraints:

1 <= nums.length <= 2000
-106 <= nums[i] <= 106
*/

// TC - O(N^2), SC - O(N)
func findNumberOfLIS(A []int) int {
	DP := make(map[int][2]int) // key = index, value = [length of LIS, count]
	N := len(A)
	LIS, LISCnt := 0, 0 // length of LIS, count of LIS

	// i = start of subseq
	for i := N - 1; i >= 0; i-- {
		maxLen, maxCnt := 1, 1 // len, cnt of LIS start from i

		for j := i + 1; j < N; j++ {
			if A[j] > A[i] {
				length, count := DP[j][0], DP[j][1] // len, cnt of LIS start from j
				if length+1 > maxLen {
					maxLen, maxCnt = length+1, count
				} else if length+1 == maxLen {
					maxCnt += count
				}
			}
		}

		if maxLen > LIS {
			LIS, LISCnt = maxLen, maxCnt
		} else if maxLen == LIS {
			LISCnt += maxCnt
		}

		DP[i] = [2]int{maxLen, maxCnt}
	}

	return LISCnt
}

func mainNL() {
	nums := []int{1, 3, 5, 4, 7}
	count := findNumberOfLIS(nums)
	fmt.Println("Number of LIS:", count)

	nums = []int{2, 2, 2, 2, 2}
	count = findNumberOfLIS(nums)
	fmt.Println("Number of LIS:", count)
}
