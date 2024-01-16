package main

import "fmt"

/**
695. Max Area of Island

You are given an m x n binary matrix grid. An island is a group of 1's (representing land)
connected 4-directionally (horizontal or vertical.) You may assume all four edges of the
grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.

Example 1:

Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0]
,[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
Example 2:

Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] is either 0 or 1.
*/

func dfsMaxArea(R, C int, grid [][]int) int {
	if R < 0 || R >= len(grid) || C < 0 || C >= len(grid[0]) || grid[R][C] == 0 {
		return 0
	}

	grid[R][C] = 0
	count := 1
	count += dfsMaxArea(R+1, C, grid)
	count += dfsMaxArea(R-1, C, grid)
	count += dfsMaxArea(R, C+1, grid)
	count += dfsMaxArea(R, C-1, grid)
	return count

}

// return dfs(R+1, C, grid) + dfs(R-1, C, grid) + dfs(R, C+1, grid)
// + dfs(R, C-1, grid) + 1

// TC - O(N*M), SC - O(N*M)
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	maxCount := 0
	for R := 0; R < len(grid); R++ {
		for C := 0; C < len(grid[0]); C++ {
			if grid[R][C] == 1 {
				maxCount = max(maxCount, dfsMaxArea(R, C, grid))
			}
		}
	}

	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func mainMaxArea() {
	grid := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println("Max area of island:", maxAreaOfIsland(grid))
}
