package main

import (
	"fmt"
	"sort"
)

// Generate all subsets of the given set of numbers without duplicates
func subsetsWithoutDuplicates(nums []int) [][]int {
	var subsets [][]int
	var curSet []int
	helper(0, nums, &curSet, &subsets)
	return subsets
}

func helper(i int, nums []int, curSet *[]int, subsets *[][]int) {
	if i >= len(nums) {
		*subsets = append(*subsets, append([]int(nil), *curSet...))
		return
	}

	// decision to include nums[i]
	*curSet = append(*curSet, nums[i])
	helper(i+1, nums, curSet, subsets)
	*curSet = (*curSet)[:len(*curSet)-1]

	// decision NOT to include nums[i]
	helper(i+1, nums, curSet, subsets)
}

// Generate all subsets of the given set of numbers with duplicates
func subsetsWithDuplicates(nums []int) [][]int {
	sort.Ints(nums)
	var subsets [][]int
	var curSet []int
	helper2(0, nums, &curSet, &subsets)
	return subsets
}

func helper2(i int, nums []int, curSet *[]int, subsets *[][]int) {
	if i >= len(nums) {
		*subsets = append(*subsets, append([]int(nil), *curSet...))
		return
	}

	// decision to include nums[i]
	*curSet = append(*curSet, nums[i])
	helper2(i+1, nums, curSet, subsets)
	*curSet = (*curSet)[:len(*curSet)-1]

	// decision NOT to include nums[i]
	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}
	helper2(i+1, nums, curSet, subsets)
}

func mainSub() {
	// nums := []int{1, 2, 3}
	// fmt.Println("Subsets without duplicates:", subsetsWithoutDuplicates(nums))
	// fmt.Println("Subsets with duplicates:", subsetsWithDuplicates(nums))
	dupNums := []int{1, 2, 2, 3} // it can be 1,2,2,2,3
	// fmt.Println("Subsets without duplicates:", subsetsWithoutDuplicates(dupNums))
	fmt.Println("Subsets with duplicates:", subsetsWithDuplicates(dupNums))
}
