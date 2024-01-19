nextRow -1
nextCol -1
nextRow -1
nextCol 0
nextRow -1
nextCol 1
nextRow 0
nextCol -1
nextRow 0
nextCol 1
queue [[0 1]]
grid [[1 2 0] [1 1 0] [1 1 0]]
nextRow 1
nextCol -1
nextRow 1
nextCol 0
nextRow 1
nextCol 1
nextRow -1
nextCol 0
nextRow -1
nextCol 1
nextRow -1
nextCol 2
nextRow 0
nextCol 0
nextRow 0
nextCol 2
queue [[0 2]]
grid [[1 2 3] [1 1 0] [1 1 0]]
nextRow 1
nextCol 0
nextRow 1
nextCol 1
nextRow 1
nextCol 2
queue [[0 2] [1 2]]
grid [[1 2 3] [1 1 3] [1 1 0]]
nextRow -1
nextCol 1
nextRow -1
nextCol 2
nextRow -1
nextCol 3
nextRow 0
nextCol 1
nextRow 0
nextCol 3
nextRow 1
nextCol 1
nextRow 1
nextCol 2
nextRow 1
nextCol 3
nextRow 0
nextCol 1
nextRow 0
nextCol 2
nextRow 0
nextCol 3
nextRow 1
nextCol 1
nextRow 1
nextCol 3
nextRow 2
nextCol 1
nextRow 2
nextCol 2
queue [[2 2]]
grid [[1 2 3] [1 1 3] [1 1 4]]
nextRow 2
nextCol 3
Shortest Path: 4


package main

import (
	"fmt"
)

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
			fmt.Println("nextRow", nextRow)
			fmt.Println("nextCol", nextCol)
			if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS || grid[nextRow][nextCol] != 0 {
				continue
			}
			queue = append(queue, [2]int{nextRow, nextCol})
			fmt.Println("queue", queue)
			grid[nextRow][nextCol] = dist + 1
			fmt.Println("grid", grid)
		}
	}

	return -1
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println("Shortest Path:", shortestPathBinaryMatrix(grid))
}
