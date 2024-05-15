package main

import (
	"fmt"
	"sort"
)

/*
Let A be an array containing integer values. The distance of A is defined as the
"minimum number"of elements in A that must be replaced with another integer so that
the resulting array is sorted in non-decreasing order. The distance of the array [2,5, 3, 1, 4, 2, 6]
*/

/**
brute force:
1. Create a sorted copy of the original array
2. Count mismatches between original and sorted arrays
3. this wont work since it is not taing care of duplicates
4. DP approach possible for this
*/

/**
Note: In the context of the "Minimum Number of Removals to Make Array Increasing" problem,
"replaced with another integer" means that you can swap the current element with any other
integer, regardless of whether that integer already exists in the array.
*/

// TC - O(NLogN), SC - O(N)
func distance(A []int) int {
	N := len(A)

	if N <= 1 {
		return 0
	}

	tail := make([]int, 1) // Initialize tail array
	tail[0] = A[0]         // First element forms an LIS of length 1
	LIS := 1               // Length of the longest LIS so far

	for i := 1; i < N; i++ {
		if A[i] > tail[LIS-1] {
			// New longest LIS, extend 'tail'
			tail = append(tail, A[i])
			LIS++
		} else {
			// Replace to maintain 'tail' sorted
			index := sort.SearchInts(tail, A[i])
			tail[index] = A[i]
		}
	}

	return N - LIS
}

func mainME() {
	nums := []int{2, 5, 3, 1, 4, 2, 6}
	minOps := distance(nums)
	fmt.Println("Minimum Operations:", minOps)
}
