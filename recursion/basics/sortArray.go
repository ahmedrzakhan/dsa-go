package main

import (
	"fmt"
)

// sort recursively sorts the array except for the last element.
// It then inserts the last element in its correct position.
func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Remove the last element and sort the rest of the array
	temp := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	arr = sort(arr)

	// Insert the removed element back into the sorted array
	return insert(arr, temp)
}

// insert places an element into its correct position in a sorted array.
func insert(arr []int, temp int) []int {
	if len(arr) == 0 || temp >= arr[len(arr)-1] {
		return append(arr, temp)
	}

	// Remove the last element, insert temp, and then add the removed element back
	val := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	arr = insert(arr, temp)

	return append(arr, val)
}

func mainSort() {
	arr := []int{5, 1, 0, 2}
	fmt.Println("Original array:", arr)
	sortedArr := sort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
