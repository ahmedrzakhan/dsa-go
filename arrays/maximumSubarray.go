package main

import "fmt"

/**
53. Maximum Subarray

Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104

Follow up: If you have figured out the O(n) solution, try coding another solution using the
divide and conquer approach, which is more subtle.
*/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// TC - O(N), SC - O(1)
func maxSubArray(arr []int) int {
	sum, maxSum := 0, 0
	for _, ele := range arr {
		sum = sum + ele
		maxSum = max(sum, maxSum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func mainMa() {
	// arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(arr))
}
