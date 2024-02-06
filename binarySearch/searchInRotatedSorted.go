package main

import "fmt"

/**
33. Search in Rotated Sorted Array

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums,
or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104
*/

// TC - O(LogN), SC - O(1)
func searchSR(arr []int, target int) int {
	L, R := 0, len(arr)-1
	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] == target {
			return mid
		}

		// Determine the sorted part of the array
		if arr[mid] >= arr[L] {
			if target >= arr[L] && target < arr[mid] {
				R = mid - 1
			} else {
				L = mid + 1
			}
		} else {
			if target > arr[mid] && target <= arr[R] {
				L = mid + 1
			} else {
				R = mid - 1
			}
		}
	}
	return -1 // Target is not found
}

func mainSR() {
	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	target := 5
	result := search(arr, target)
	fmt.Println("Index of target is:", result) // Output: Index of target is: 4

	arr = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	result = search(arr, target)
	fmt.Println("Index of target is:", result) // Output: Index of target is: 4

	// target = 3
	// result = search(arr, target)
	// fmt.Println("Index of target is:", result) // Output: Index of target is: -1

	arr = []int{6, 7, 0, 1, 2, 4, 5}
	target = 5
	result = search(arr, target)
	fmt.Println("Index of target is:", result) // Output: Index of target is: -1

}
