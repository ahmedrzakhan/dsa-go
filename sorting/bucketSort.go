package main

import "fmt"

func bucketSort(arr []int) []int {
	// Assuming arr only contains 0, 1 or 2
	counts := [3]int{0, 0, 0}

	// Count the quantity of each val in arr
	for _, num := range arr {
		counts[num]++
	}

	// Fill each bucket in the original array
	i := 0
	for idx, val := range counts {
		for j := 0; j < val; j++ {
			arr[i] = idx
			i++
		}
	}
	return arr
}

func mainBS() {
	arr := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(bucketSort(arr))
}
