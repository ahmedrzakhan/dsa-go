package main

import (
	"fmt"
	"math/rand"
)

/**

912. Sort an Array

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity
 and with the smallest space complexity possible.

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

// partitionDesc rearranges the elements in the array, so that all elements greater than
// the pivot are on the left of the pivot, and all elements lesser are on the right.
func partitionDesc(arr []int, L, R int) int {
	pivot := arr[R]
	i := L

	for j := L; j < R; j++ {
		if arr[j] > pivot { // Changed from < to >
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[R] = arr[R], arr[i]
	return i
}

// partition rearranges the elements in the array, so that all elements less than
// the pivot are on the left of the pivot, and all elements greater are on the right.
func partition(arr []int, L, R int) int {

	// Randomizing the pivot
	pivotIndex := rand.Intn(R-L+1) + L
	arr[pivotIndex], arr[R] = arr[R], arr[pivotIndex]

	pivot := arr[R]
	i := L

	for j := L; j < R; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[R] = arr[R], arr[i]
	return i
}

// QuickSort sorts an array using the QuickSort algorithm.
func QuickSort(arr []int, L, R int) {
	if L < R {
		pivot := partition(arr, L, R)
		QuickSort(arr, L, pivot-1)
		QuickSort(arr, pivot+1, R)
	}
}

// TC - O(NLOGN)/ O(N^2), SC - O(1)
func QuickSortDesc(arr []int, L, R int) {
	if L < R {
		pivot := partitionDesc(arr, L, R)
		QuickSortDesc(arr, L, pivot-1)
		QuickSortDesc(arr, pivot+1, R)
	}
}

func mainQuick() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Original array:", arr)

	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:  ", arr)

	QuickSortDesc(arr, 0, len(arr)-1)
	fmt.Println("Sorted array desc:  ", arr)
}
