package main

import (
	"fmt"
	"math"
)

/**
410. Split Array Largest Sum

Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.

Return the minimized largest sum of the split.

A subarray is a contiguous part of the array.

Example 1:

Input: nums = [7,2,5,10,8], k = 2
Output: 18
Explanation: There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8], where the largest sum among the two subarrays is only 18.
Example 2:

Input: nums = [1,2,3,4,5], k = 2
Output: 9
Explanation: There are four ways to split nums into two subarrays.
The best way is to split it into [1,2,3] and [4,5], where the largest sum among the two subarrays is only 9.

Constraints:

1 <= nums.length <= 1000
0 <= nums[i] <= 106
1 <= k <= min(50, nums.length)
*/

// TC - O(N^2*K), SC - O(1)
func splitArray(arr []int, K int) int {
	return split(arr, 0, K)
}

// Brute force
func split(arr []int, index, K int) int {
	if K == 1 {
		return sum(arr[index:])
	}
	minLargestSum := math.MaxInt
	for i := index; i <= len(arr)-K; i++ {
		firstPartSum := sum(arr[index : i+1])
		largestSum := max(firstPartSum, split(arr, i+1, K-1))
		minLargestSum = min(minLargestSum, largestSum)
		if firstPartSum > minLargestSum {
			break
		}
	}
	return minLargestSum
}

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func max(a, b int) int {
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

// TC - O(NLog(S)), SC - O(1) S is sum
func splitArrayOp(arr []int, K int) int {
	L, R := 0, 0
	for _, num := range arr {
		R += num // R is sum
		if L < num {
			L = num // L is largest value in array
		}
	}

	result := R
	for L <= R {
		mid := L + (R-L)/2
		if isValid(arr, K, mid) {
			result = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	return result
}

func isValid(arr []int, K, maxSum int) bool {
	count, sum := 1, 0
	for _, val := range arr {
		sum += val
		if sum > maxSum {
			count++
			sum = val
			if count > K {
				return false
			}
		}
	}
	return true
}

func mainSA() {
	arr := []int{7, 2, 5, 10, 8}
	K := 3
	fmt.Println("Minimum largest split sum:", splitArray(arr, K))
}
