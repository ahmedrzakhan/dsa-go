package main

import "fmt"

/**
62. Unique Paths

There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either
down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach
the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:

Input: m = 3, n = 7
Output: 28
Example 2:

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

Constraints:

1 <= m, n <= 100
*/

// Backtracking function
func helper(R, C int, count *int, ROWS, COLS int) {
	if R == ROWS-1 && C == COLS-1 { // Reached the target
		*count++
		return
	}

	// Explore moving right
	if C < COLS-1 {
		helper(R, C+1, count, ROWS, COLS)
	}

	// Explore moving down
	if R < ROWS-1 {
		helper(R+1, C, count, ROWS, COLS)
	}
}

// TC - O(2^(R+C)), SC - O(R+C)
func uniquePaths(ROWS, COLS int) int {
	count := 0                       // Reset count for each calculation
	helper(0, 0, &count, ROWS, COLS) // Start from the top-left
	return count
}

// TC - O(R*C), SC - O(R)
func uniquePathsDP(ROWS, COLS int) int {
	prevRow := make([]int, COLS)

	for R := ROWS - 1; R >= 0; R-- {
		curRow := make([]int, COLS)
		curRow[COLS-1] = 1 // Base case: Rightmost cell has 1 path

		for C := COLS - 2; C >= 0; C-- {
			curRow[C] = curRow[C+1] + prevRow[C]
		}

		prevRow = curRow
	}
	return prevRow[0]
}

func mainUP() {
	R := 2
	C := 2
	fmt.Println(uniquePaths(R, C)) // Output: 2

	R = 3
	C = 3
	fmt.Println(uniquePaths(R, C)) // Output: 6

	R = 5
	C = 5
	fmt.Println(uniquePaths(R, C)) // Output: 70
}
