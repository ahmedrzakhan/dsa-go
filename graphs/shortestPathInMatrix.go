package main

import (
	"fmt"
)

/**
1091. Shortest Path in Binary Matrix

Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If
there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right
cell (i.e., (n - 1, n - 1)) such that:

All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they
share an edge or a corner).
The length of a clear path is the number of visited cells of this path.

Example 1:

Input: grid = [[0,1],[1,0]]
Output: 2
Example 2:

Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
Output: 4
Example 3:

Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
Output: -1

Constraints:

n == grid.length
n == grid[i].length
1 <= n <= 100
grid[i][j] is 0 or 1
*/

// TC - O(M*N), SC - O(min(M, N))
func shortestPathBinaryMatrix(grid [][]int) int {
	ROWS, COLS := len(grid), len(grid[0])
	if grid[0][0] != 0 || grid[ROWS-1][COLS-1] != 0 {
		return -1
	}

	queue := [][2]int{{0, 0}}
	grid[0][0] = 1

	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		R, C, dist := cell[0], cell[1], grid[cell[0]][cell[1]]

		if R == ROWS-1 && C == COLS-1 {
			return dist
		}

		for _, dir := range directions {
			nextRow, nextCol := R+dir[0], C+dir[1]
			if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS || grid[nextRow][nextCol] != 0 {
				continue
			}
			queue = append(queue, [2]int{nextRow, nextCol})
			grid[nextRow][nextCol] = dist + 1
		}
	}

	return -1
}

func mainShortest() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println("Shortest Path:", shortestPathBinaryMatrix(grid))
}
