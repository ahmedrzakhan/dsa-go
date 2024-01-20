package main

import "fmt"

/**
994. Rotting Oranges

You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange
becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1.



Example 1:


Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
Example 2:

Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten,
because rotting only happens 4-directionally.
Example 3:

Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] is 0, 1, or 2.
*/

// TC - O(M*N), SC - O(M*N)
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	ROWS, COLS := len(grid), len(grid[0])
	count := 0
	queue := make([][2]int, 0)

	// First pass to count fresh oranges and enqueue rotten oranges.
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	// BFS
	for len(queue) > 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			cell := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				nextRow, nextCol := cell[0]+dir[0], cell[1]+dir[1]
				if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS {
					continue
				}
				if grid[nextRow][nextCol] == 1 {
					grid[nextRow][nextCol] = 2
					count--
					queue = append(queue, [2]int{nextRow, nextCol})
				}
			}
		}
		if len(queue) > 0 {
			minutes++
		}
	}

	if count == 0 {
		return minutes
	}
	return -1
}

func mainRot() {
	// grid := [][]int{
	// 	{2, 1, 1},
	// 	{1, 1, 0},
	// 	{0, 1, 1},
	// }
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 2},
	}
	result := orangesRotting(grid)
	fmt.Printf("Minutes until all oranges rot: %d\n", result)
}
