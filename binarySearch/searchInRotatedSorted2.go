package main

import "fmt"

/**
81. Search in Rotated Sorted Array II

There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
(0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if
it is not in nums.

You must decrease the overall operation steps as much as possible.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums is guaranteed to be rotated at some pivot.
-104 <= target <= 104

Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?
*/

// TC - O(N), SC - O(1)
func searchSRA(arr []int, target int) bool {
	L, R := 0, len(arr)-1

	for L <= R {
		mid := L + (R-L)/2

		if arr[mid] == target {
			return true
		}

		// If we have duplicates, we just shrink the range
		if arr[L] == arr[mid] && arr[R] == arr[mid] {
			L++
			R--
		} else if arr[mid] >= arr[L] { // L side is sorted
			if target >= arr[L] && target < arr[mid] {
				R = mid - 1
			} else {
				L = mid + 1
			}
		} else { // R side is sorted
			if target > arr[mid] && target <= arr[R] {
				L = mid + 1
			} else {
				R = mid - 1
			}
		}
	}

	return false
}

func mainSRA() {
	arr := []int{3, 1}
	target := 1
	result := search(arr, target)
	fmt.Println("found:", result) // true
	// arr := []int{2, 5, 6, 0, 0, 1, 2}
	// target := 0
	// result := search(arr, target)
	// fmt.Println("found:", result) // true

	arr = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	result = search(arr, target)
	fmt.Println("found:", result) // false

	arr = []int{1, 0, 1, 1, 1}
	target = 0
	result = search(arr, target)
	fmt.Println("found:", result) // true
}
