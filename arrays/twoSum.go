package main

import "fmt"

/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

/**
variation - return true if there exists a pair in arr whose sum = target
arr := []int{1, 3, 2, 4, 5}
target := 6

brute force: try out all possibilities, TC - O(N^2), SC - O(1)
optimized: sort and use two pointer, TC - O(NLogN), SC - O(1)
more optimized: maintain a hashmap with index as key and element as value,
iterate over array if target - curElem exists in hashmap, return, TC - O(N), SC - O(N)
*/

// TC - O(N), SC - O(N)
func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)

	for i, num := range nums {
		if idx, found := idxMap[target-num]; found {
			return []int{idx, i}
		}

		idxMap[num] = i
	}

	return nil
}

func mainTwoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
