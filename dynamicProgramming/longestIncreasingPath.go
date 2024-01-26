package main

import "fmt"

/**
329. Longest Increasing Path in a Matrix

Given an m x n integers matrix, return the length of the longest increasing path in matrix.

From each cell, you can either move in four directions: left, right, up, or down. You may not
move diagonally or move outside the boundary (i.e., wrap-around is not allowed).

Example 1:

Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:

Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
Example 3:

Input: matrix = [[1]]
Output: 1

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
0 <= matrix[i][j] <= 231 - 1
*/

// TC - O(R*C), SC - O(R*C)
func longestIncreasingPath(matrix [][]int) int {
	ROWS, COLS := len(matrix), len(matrix[0])

	if ROWS == 0 || COLS == 0 {
		return 0
	}

	dpMatrix := make([][]int, ROWS)
	for i := range dpMatrix {
		dpMatrix[i] = make([]int, COLS)
	}

	maxPath := 0
	for R := 0; R < ROWS; R++ {
		for C := 0; C < COLS; C++ {
			maxPath = max(maxPath, dfsLIP(R, C, ROWS, COLS, matrix, dpMatrix))
		}
	}
	return maxPath
}

func dfsLIP(R, C, ROWS, COLS int, matrix, dpMatrix [][]int) int {
	if dpMatrix[R][C] != 0 {
		return dpMatrix[R][C]
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range directions {
		nextRow, nextCol := R+dir[0], C+dir[1]

		if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS || matrix[R][C] >= matrix[nextRow][nextCol] {
			continue
		}

		dpMatrix[R][C] = max(dpMatrix[R][C], dfsLIP(nextRow, nextCol, ROWS, COLS, matrix, dpMatrix))
	}

	dpMatrix[R][C]++
	return dpMatrix[R][C]
}

func maxLIP(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mainLIP() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	fmt.Println("Longest Increasing Path:", longestIncreasingPath(matrix))

	matrix2 := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	fmt.Println("Longest Increasing Path:", longestIncreasingPath(matrix2))
}
