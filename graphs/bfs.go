package main

import (
	"fmt"
)

// TC - O(N*M), SC - O(N*M)
func bfs(grid [][]int) int {
	ROWS, COLS := len(grid), len(grid[0])
	visited := make([][]int, ROWS)
	for i := range visited {
		visited[i] = make([]int, COLS)
	}

	queue := [][]int{{0, 0}} // Initialize queue with starting position
	visited[0][0] = 1

	length := 0
	for len(queue) > 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			cell := queue[0]
			queue = queue[1:] // Dequeue
			R, C := cell[0], cell[1]

			if R == ROWS-1 && C == COLS-1 {
				return length
			}

			// Add valid neighbors to queue
			directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, dir := range directions {
				nextRow, nextCol := R+dir[0], C+dir[1]
				if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS || visited[nextRow][nextCol] == 1 || grid[nextRow][nextCol] == 1 {
					continue
				}
				queue = append(queue, []int{nextRow, nextCol})
				visited[nextRow][nextCol] = 1
			}
		}
		length++
	}
	return length
}

func mainBfs() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	fmt.Println("Shortest path length:", bfs(grid))
}
