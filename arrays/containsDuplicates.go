package main

import "fmt"

/**
217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

// TC - O(N), SC - O(N)
func containsDuplicates(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}

		seen[num] = struct{}{}
	}
	return false
}

func mainContainsDuplicates() {
	fmt.Println(containsDuplicates([]int{1, 2, 3, 1}))
}

// Approach 2 - SORT and iterate if next elem is same it contains duplicate
// TC - O(N) SC - O(1)
