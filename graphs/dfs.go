package main

import "fmt"

func dfs(R, C, ROWS, COLS int, visited map[[2]int]bool, grid [][]int) int {
	if R < 0 || C < 0 || R >= ROWS || C >= COLS || visited[[2]int{R, C}] || grid[R][C] == 1 {
		return 0
	}
	if R == ROWS-1 && C == COLS-1 {
		return 1
	}

	visited[[2]int{R, C}] = true

	count := 0
	count += dfs(R+1, C, ROWS, COLS, visited, grid)
	count += dfs(R-1, C, ROWS, COLS, visited, grid)
	count += dfs(R, C+1, ROWS, COLS, visited, grid)
	count += dfs(R, C-1, ROWS, COLS, visited, grid)

	delete(visited, [2]int{R, C})
	return count
}

func countPaths(grid [][]int) int {
	ROWS, COLS := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	return dfs(0, 0, ROWS, COLS, visited, grid)
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	fmt.Println("Number of paths:", countPaths(grid))
}
