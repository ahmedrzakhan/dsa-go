package main

import (
	"fmt"
)

/**
78. Subsets

Given an integer array nums of unique elements, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/

// TC - O(N*2^N), SC - O(N*2^N)
func subsets(nums []int) [][]int {
	var subsets [][]int
	var curSet []int
	helper(0, nums, &curSet, &subsets)
	return subsets
}

func helpers(i int, nums []int, curSet *[]int, subsets *[][]int) {
	if i >= len(nums) {
		*subsets = append(*subsets, append([]int(nil), *curSet...))
		return
	}

	*curSet = append(*curSet, nums[i])
	helper(i+1, nums, curSet, subsets)
	*curSet = (*curSet)[:len(*curSet)-1]

	helper(i+1, nums, curSet, subsets)
}

func mainSu() {
	dupNums := []int{1, 2, 2, 3} // it can be 1,2,2,2,3
	fmt.Println("Subsets with duplicates:", subsets(dupNums))
}
