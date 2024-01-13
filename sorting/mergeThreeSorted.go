package main

import (
	"fmt"
	"math"
)

/**
Merge three sorted arrays
*/
// TC - O(N), SC - O(1)
func mergeThree(arr1, arr2, arr3 []int) []int {
	i, j, k := 0, 0, 0
	N1, N2, N3 := len(arr1), len(arr2), len(arr3)
	result := make([]int, 0, N1+N2+N3)

	for i < N1 || j < N2 || k < N3 {
		minVal := math.MaxInt
		if i < N1 && arr1[i] < minVal {
			minVal = arr1[i]
		}
		if j < N2 && arr2[j] < minVal {
			minVal = arr2[j]
		}
		if k < N3 && arr3[k] < minVal {
			minVal = arr3[k]
		}

		result = append(result, minVal)

		if i < N1 && arr1[i] == minVal {
			i++
		} else if j < N2 && arr2[j] == minVal {
			j++
		} else if k < N3 && arr3[k] == minVal {
			k++
		}
	}

	return result
}

func mainMergeThree() {
	arr1 := []int{1, 2, 3, 4, 9, 10}
	arr2 := []int{1, 2, 3, 4, 5, 6, 11, 12}
	arr3 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(mergeThree(arr1, arr2, arr3))
}
