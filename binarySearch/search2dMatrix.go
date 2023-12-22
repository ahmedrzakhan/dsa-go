package main

import "fmt"

/**
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/
// TC - O(log(m*n)), SC - O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	ROWS := len(matrix)
	COLS := len(matrix[0])

	top := 0
	bot := ROWS - 1

	for top <= bot {
		row := (top + bot) / 2
		if target > matrix[row][len(matrix[row])-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bot = row - 1
		} else {
			break
		}
	}

	if !(top <= bot) {
		return false
	}

	row := (top + bot) / 2
	left := 0
	right := COLS - 1
	for left <= right {
		mid := (left + right) / 2
		if target > matrix[row][mid] {
			left = mid + 1
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func mainSear() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	target := 3
	found := searchMatrix(matrix, target)
	fmt.Printf("Found %d in matrix: %v\n", target, found)
}
