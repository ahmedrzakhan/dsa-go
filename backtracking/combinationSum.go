package main

import (
	"fmt"
)

/**
39. Combination Sum

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of
candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency
 of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations
for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []
*/

func combinationSum(arr []int, target int) [][]int {
	subsets := make([][]int, 0)
	curSet := make([]int, 0)
	helperComb(0, arr, target, 0, &curSet, &subsets)
	return subsets
}

// TC - O(2^N+M), SC - O(N * M)
func helperComb(idx int, arr []int, target int, curSum int, curSet *[]int, subsets *[][]int) {
	if curSum == target {
		*subsets = append(*subsets, append([]int{}, *curSet...))
		return
	}
	if curSum > target {
		return
	}
	for i := idx; i < len(arr); i++ {
		*curSet = append(*curSet, arr[i])
		helperComb(i, arr, target, curSum+arr[i], curSet, subsets)
		*curSet = (*curSet)[:len(*curSet)-1]
	}
}

func mainCom() {
	arr := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(arr, target)
	fmt.Println("Combinations that add up to", target, "are:", result)
}
