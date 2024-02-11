package main

import "fmt"

/**
219. Contains Duplicate II

Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/

// TC - O(N), SC - O(N)
func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, ele := range nums {
		if prevIndex, exists := indexMap[ele]; exists {
			if i-prevIndex <= k {
				return true
			}
		}
		indexMap[ele] = i
	}
	return false
}

func mainCon() {
	arr := []int{1, 2, 3, 1}
	k := 3

	fmt.Println(containsNearbyDuplicate(arr, k))
}
