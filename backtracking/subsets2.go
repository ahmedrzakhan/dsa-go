package main

import (
	"fmt"
	"sort"
)

/**
90. Subsets II

Given an integer array nums that may contain duplicates, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10

*/

// TC - O(N*2^N),SC - O(N)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var subsets [][]int
	var curSet []int
	helper(0, nums, &curSet, &subsets)
	return subsets
}

func helperss(i int, nums []int, curSet *[]int, subsets *[][]int) {
	if i >= len(nums) {
		*subsets = append(*subsets, append([]int(nil), *curSet...))
		return
	}

	*curSet = append(*curSet, nums[i])
	helper(i+1, nums, curSet, subsets)
	*curSet = (*curSet)[:len(*curSet)-1]

	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}
	helper(i+1, nums, curSet, subsets)
}

func mainSubs() {
	// nums := []int{1, 2, 3}
	// fmt.Println("Subsets without duplicates:", subsetsWithoutDuplicates(nums))
	// fmt.Println("Subsets with duplicates:", subsetsWithDuplicates(nums))
	dupNums := []int{1, 2, 2, 3} // it can be 1,2,2,2,3
	// fmt.Println("Subsets without duplicates:", subsetsWithoutDuplicates(dupNums))
	fmt.Println("Subsets with duplicates:", subsetsWithDuplicates(dupNums))
}
