package main

import "fmt"

func countLargerthanNeighbour(arr []int) int {
	count := 0
	for i := 1; i < len(arr)-2; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	result := countLargerthanNeighbour(arr)
	fmt.Println(result)
}
