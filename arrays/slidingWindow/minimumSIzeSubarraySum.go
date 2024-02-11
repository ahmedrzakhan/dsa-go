package main

import "fmt"

/**
209. Minimum Size Subarray Sum

Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0


Constraints:

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104


Follow up: If you have figured out the O(n) solution, try coding another solution of which the time
complexity is O(n log(n)).
*/

/**
Note: cannot use hashmap here since subarray needs to be greater equal to not exact target
like in longest subarray with sum k
*/

// TC - O(N), SC - O(1)
func minSubArrayLen(target int, arr []int) int {
	N := len(arr)
	minLength := N + 1 // Initialize with a max value greater than the possible max length
	sum := 0
	L := 0

	for R := 0; R < N; R++ {
		sum += arr[R]
		// Shrink the window from the L as long as the sum is greater than or equal to target
		for sum >= target {
			if R-L+1 < minLength {
				minLength = R - L + 1
			}
			sum -= arr[L]
			L++
		}
	}

	if minLength == N+1 {
		return 0 // No valid subarray found
	}
	return minLength
}

func mainMS() {
	target := 7
	arr := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, arr))
}
