package main

import "fmt"

/**
912. Sort an Array

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time
complexity and with the smallest space complexity possible.

Example 1:

Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not changed (for example,
2 and 3), while the positions of other numbers are changed (for example, 1 and 5).
Example 2:

Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessairly unique.

Constraints:

1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104
*/

/**
NOTE: For preserving relative order of elements it should be
leftArr[i] <= rightArr[j]
since merge sort is a stable sorting algorithm
also
result = append(result, leftArr[i:]...)
result = append(result, rightArr[j:]...)
first left than right
*/

// merge merges two sorted slices into a single sorted slice.
func merge(leftArr, rightArr []int) []int {
	result := make([]int, 0, len(leftArr)+len(rightArr))

	i, j := 0, 0

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			result = append(result, leftArr[i])
			i++
		} else {
			result = append(result, rightArr[j])
			j++
		}
	}

	// Append any remaining elements
	result = append(result, leftArr[i:]...)
	result = append(result, rightArr[j:]...)

	return result
}

// MergeSort performs the merge sort algorithm on a slice of integers.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	leftArr := MergeSort(arr[:mid])
	rightArr := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(leftArr, rightArr)
}

func mainMerge() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)
	sortedArr := MergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
