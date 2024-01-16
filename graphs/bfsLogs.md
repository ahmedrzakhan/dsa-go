{0!, 0!, 0!, 0!},
{1, 1, 0, 0},
{0, 0, 0, 1},
{0, 1, 0, 0},

queueLength 1
queue dequeue []
R, C 0 0
nextRow, nextCol 1 0
nextRow, nextCol -1 0
nextRow, nextCol 0 1
queue append [[0 1]]
nextRow, nextCol 0 -1
length++ 1
queue [[0 1]]
queueLength 1
queue dequeue []
R, C 0 1
nextRow, nextCol 1 1
nextRow, nextCol -1 1
nextRow, nextCol 0 2
queue append [[0 2]]
nextRow, nextCol 0 0
length++ 2
queue [[0 2]]
queueLength 1
queue dequeue []
R, C 0 2
nextRow, nextCol 1 2
queue append [[1 2]]
nextRow, nextCol -1 2
nextRow, nextCol 0 3
queue append [[1 2] [0 3]]
nextRow, nextCol 0 1
length++ 3
queue [[1 2] [0 3]]
queueLength 2
queue dequeue [[0 3]]
R, C 1 2
nextRow, nextCol 2 2
queue append [[0 3] [2 2]]
nextRow, nextCol 0 2
nextRow, nextCol 1 3
queue append [[0 3] [2 2] [1 3]]
nextRow, nextCol 1 1
queue dequeue [[2 2] [1 3]]
R, C 0 3
nextRow, nextCol 1 3
nextRow, nextCol -1 3
nextRow, nextCol 0 4
nextRow, nextCol 0 2
length++ 4
queue [[2 2] [1 3]]
queueLength 2
queue dequeue [[1 3]]
R, C 2 2
nextRow, nextCol 3 2
queue append [[1 3] [3 2]]
nextRow, nextCol 1 2
nextRow, nextCol 2 3
nextRow, nextCol 2 1
queue append [[1 3] [3 2] [2 1]]
queue dequeue [[3 2] [2 1]]
R, C 1 3
nextRow, nextCol 2 3
nextRow, nextCol 0 3
nextRow, nextCol 1 4
nextRow, nextCol 1 2
length++ 5
queue [[3 2] [2 1]]
queueLength 2
queue dequeue [[2 1]]
R, C 3 2
nextRow, nextCol 4 2
nextRow, nextCol 2 2
nextRow, nextCol 3 3
queue append [[2 1] [3 3]]
nextRow, nextCol 3 1
queue dequeue [[3 3]]
R, C 2 1
nextRow, nextCol 3 1
nextRow, nextCol 1 1
nextRow, nextCol 2 2
nextRow, nextCol 2 0
queue append [[3 3] [2 0]]
length++ 6
queue [[3 3] [2 0]]
queueLength 2
queue dequeue [[2 0]]
R, C 3 3
Shortest path length: 6

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
		fmt.Println("queueLength", queueLength)
		for i := 0; i < queueLength; i++ {
			cell := queue[0]
			queue = queue[1:] // Dequeue
			fmt.Println("queue dequeue", queue)
			R, C := cell[0], cell[1]
			fmt.Println("R, C", R, C)

			if R == ROWS-1 && C == COLS-1 {
				return length
			}

			// Add valid neighbors to queue
			directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, dir := range directions {
				nextRow, nextCol := R+dir[0], C+dir[1]
				fmt.Println("nextRow, nextCol", nextRow, nextCol)
				if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS || visited[nextRow][nextCol] == 1 || grid[nextRow][nextCol] == 1 {
					continue
				}
				queue = append(queue, []int{nextRow, nextCol})
				fmt.Println("queue append", queue)
				visited[nextRow][nextCol] = 1
			}
		}
		length++
		fmt.Println("length++", length)
		fmt.Println("queue", queue)
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

