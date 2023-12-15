package main

import (
	"fmt"
)

/**
978. Longest Turbulent Subarray

Given an integer array arr, return the length of a maximum size turbulent subarray of arr.

A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent if and only if:

For i <= k < j:
arr[k] > arr[k + 1] when k is odd, and
arr[k] < arr[k + 1] when k is even.
Or, for i <= k < j:
arr[k] > arr[k + 1] when k is even, and
arr[k] < arr[k + 1] when k is odd.


Example 1:

Input: arr = [9,4,2,10,7,8,8,1,9]
Output: 5
Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
Example 2:

Input: arr = [4,8,12,16]
Output: 2
Example 3:

Input: arr = [100]
Output: 1


Constraints:

1 <= arr.length <= 4 * 104
0 <= arr[i] <= 109
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TC - O(N), SC - O(1)
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n < 2 {
		return n
	}

	maxLength, inc, dec := 1, 1, 1

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			inc = dec + 1
			dec = 1
		} else if arr[i] < arr[i-1] {
			dec = inc + 1
			inc = 1
		} else {
			// if  its equal equal to like 8, 8
			inc, dec = 1, 1
		}
		maxLength = max(maxLength, max(inc, dec))
	}

	return maxLength
}

func mainLong() {
	arr := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	fmt.Println("Maximum length of a turbulent subarray:", maxTurbulenceSize(arr))
}
