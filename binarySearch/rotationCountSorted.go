package main

import "fmt"

/**
how many times sorted array is rotated
*/

func findMinIndex(A []int) int {
	L, R := 0, len(A)-1

	// Base Case: Array is already sorted (No rotation)
	if A[L] <= A[R] {
		return 0 // 0 rotations
	}

	for L <= R {
		mid := L + (R-L)/2

		// Check if mid is the minimum (point of rotation)
		if A[mid] > A[mid+1] {
			return mid + 1 // Rotations = index of minimum + 1
		}

		// Decide which half is sorted
		if A[L] <= A[mid] {
			// Left half is sorted, search the right half
			L = mid + 1
		} else {
			// Right half is sorted, search the left half
			R = mid - 1
		}
	}

	return 0 // Should not reach here if the input is a rotated sorted array
}

func mainRCS() {
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	rotations1 := findMinIndex(nums1)
	fmt.Println("Array 1 Rotations:", rotations1)

	nums2 := []int{7, 8, 1, 2, 3, 5}
	rotations2 := findMinIndex(nums2)
	fmt.Println("Array 2 Rotations:", rotations2)
}
