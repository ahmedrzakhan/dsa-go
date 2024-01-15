package main

import "fmt"

/*
200. Number of Islands

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally
or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

// TC - O(M*N), SC - O(M*N)
func dfsIsland(R, C int, grid [][]byte) {
	if R < 0 || C < 0 || R >= len(grid) || C >= len(grid[0]) || grid[R][C] == '0' {
		return
	}

	grid[R][C] = '0'
	dfsIsland(R+1, C, grid)
	dfsIsland(R-1, C, grid)
	dfsIsland(R, C+1, grid)
	dfsIsland(R, C-1, grid)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	for R := 0; R < len(grid); R++ {
		for C := 0; C < len(grid[0]); C++ {
			if grid[R][C] == '1' {
				count++
				dfsIsland(R, C, grid)
			}
		}
	}

	return count
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	// grid := [][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }
	fmt.Println("Number of islands:", numIslands(grid))
}
