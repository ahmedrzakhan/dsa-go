package main

import "fmt"

/**
Given an integer array nums of length n, you want to create an array ans of length 2n
 where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

Specifically, ans is the concatenation of two nums arrays.

Return the array ans.



Example 1:

Input: nums = [1,2,1]
Output: [1,2,1,1,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
- ans = [1,2,1,1,2,1]
Example 2:

Input: nums = [1,3,2,1]
Output: [1,3,2,1,1,3,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]
- ans = [1,3,2,1,1,3,2,1]

Constraints:

n == nums.length
1 <= n <= 1000
1 <= nums[i] <= 1000
*/

// TC - O(N) SC O(1)
func getConcatenation(nums []int) []int {
	nums = append(nums, make([]int, len(nums))...)
	N := len(nums)
	for i := 0; i < N/2; i++ {
		nums[i+N/2] = nums[i]
	}
	return nums
}

func mainConcatenation() {
	arr := []int{1, 3, 2, 1}
	// 1, 3, 2
	// check wit odd and even nums of items
	fmt.Println(getConcatenation(arr))
}
