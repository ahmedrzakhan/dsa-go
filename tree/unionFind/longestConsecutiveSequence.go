package main

import "fmt"

/**
128. Longest Consecutive Sequence

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

// TC - O(N), SC - O(N)
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	maxLen := 0
	for num := range numsMap {
		if !numsMap[num-1] { // Check if it's the start of a sequence
			length := 1
			for numsMap[num+length] { // Extend the sequence
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}

func mainLCS() {
	fmt.Println(longestConsecutive([]int{4, 7, 3, 8, 2, 1}))                // Output: 4
	fmt.Println(longestConsecutive([]int{4, 7, 3, 8, 2, 1, 9, 24, 10, 11})) // Output: 5
}
