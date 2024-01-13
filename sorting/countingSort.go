package main

import (
	"fmt"
)

/**
you can modify Counting Sort to work with negative integers as well. The key modification involves adjusting
the range of the count array to accommodate negative values. Here's how you could do it:
Find Minimum and Maximum Values: Instead of just finding the maximum value, also find the minimum value
in the array.
Adjust Count Array: Create a count array that covers the range from the minimum value to the maximum value.
This might involve shifting the indices to ensure they are non-negative.
*/

// CountingSort sorts an array using the Counting Sort algorithm.
func CountingSort(arr []int) []int {
	// Find the maximum element in the array
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	// Create a count array with size max+1
	count := make([]int, max+1)

	// Count the occurrences of each element
	for _, value := range arr {
		count[value]++
	}

	// Modify the count array by adding the current count with the previous count
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	output := make([]int, len(arr))
	for _, value := range arr {
		output[count[value]-1] = value
		count[value]--
	}

	return output
}

func mainCounting() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println("Original Array:", arr)
	sortedArr := CountingSort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
