package main

import (
	"fmt"
)

/**
198. House Robber

You are a professional robber planning to rob houses along a street. Each house has a certain amount
of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses
have security systems connected and it will automatically contact the police if two adjacent houses
were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount
of money you can rob tonight without alerting the police.

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/

// TC - O(2^N), SC - O(N)
// Brute Force
func robBruteForce(A []int) int {
	return helperRob(A, 0)
}

func helperRob(A []int, i int) int {
	if i >= len(A) {
		return 0
	}

	// Option 1: Rob the current house
	robCurrent := A[i] + helperRob(A, i+2) // Add loot and skip the next

	// Option 2: Skip the current house
	skipCurrent := helperRob(A, i+1) // Move to the next house

	return max(robCurrent, skipCurrent)
}

// TC - O(N), SC - O(N)
// Dynamic Programming
func robDP(A []int) int {
	N := len(A)

	// Create a DP table to store maximum loot up to each house
	DP := make([]int, N)

	// Base Cases:
	DP[0] = A[0] // First house
	if N >= 2 {
		DP[1] = max(A[0], A[1]) // Second house
	}

	// Fill the DP table in a bottom-up fashion
	for i := 2; i < N; i++ {
		DP[i] = max(DP[i-2]+A[i], DP[i-1])
	}

	return DP[N-1] // Maximum loot is in the last element
}

func mainRo() {
	nums := []int{1, 2, 3}
	// fmt.Println("Brute Force:", robDP(nums))
	fmt.Println("Dynamic Programming:", robDP(nums))
}
