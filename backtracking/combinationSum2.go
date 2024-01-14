package main

import (
	"fmt"
	"sort"
)

/**
40. Combination Sum II

Given a collection of candidate numbers (candidates) and a target number (target), find all
unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

// TODO: get TC and SC
// TC - O(2^N), SC - O(N)
func combinationSum2(arr []int, target int) [][]int {
	sort.Ints(arr)
	curSet := make([]int, 0)
	subsets := make([][]int, 0)

	if len(arr) == 0 {
		return subsets
	}

	helperCombSum2(0, 0, arr, target, &curSet, &subsets)
	return subsets
}

func helperCombSum2(idx, curSum int, arr []int, target int, curSet *[]int, subsets *[][]int) {
	if curSum == target {
		*subsets = append(*subsets, append([]int{}, *curSet...))
		return
	}
	if curSum > target {
		return
	}

	for i := idx; i < len(arr); i++ {
		if i > idx && arr[i] == arr[i-1] {
			continue
		}
		*curSet = append(*curSet, arr[i])
		helperCombSum2(i+1, curSum+arr[i], arr, target, curSet, subsets)
		*curSet = (*curSet)[:len(*curSet)-1]
	}
}

func mainComb() {
	arr := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(arr, target))
}
