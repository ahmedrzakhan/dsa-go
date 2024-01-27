package main

import "fmt"

/**
79. Word Search

Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are
horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
Example 2:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
Example 3:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Constraints:

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.

Follow up: Could you use search pruning to make your solution faster with a larger board?
*/

// TC - O(R*C*4^N), SC - O(N)
func exist(matrix [][]byte, word string) bool {
	ROWS, COLS := len(matrix), len(matrix[0])

	if ROWS == 0 || COLS == 0 || !charsInMatrix(matrix, word) {
		return false
	}

	// Start the search from every cell
	for R := 0; R < ROWS; R++ {
		for C := 0; C < COLS; C++ {
			if dfs(0, R, C, ROWS, COLS, word, matrix) {
				return true
			}
		}
	}

	return false
}

func dfs(i, R, C, ROWS, COLS int, word string, matrix [][]byte) bool {
	// Base cases
	if i == len(word) {
		return true
	}
	if R < 0 || C < 0 || R >= ROWS || C >= COLS || matrix[R][C] != word[i] || matrix[R][C] == '#' {
		return false
	}

	// Temporarily mark the cell as visited
	temp := matrix[R][C]
	matrix[R][C] = '#'

	// Explore all possible directions
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range directions {
		nextRow, nextCol := R+dir[0], C+dir[1]

		if dfs(i+1, nextRow, nextCol, ROWS, COLS, word, matrix) {
			return true
		}
	}

	// Backtrack and unmark the visited cell
	matrix[R][C] = temp

	return false
}

func charsInMatrix(matrix [][]byte, word string) bool {
	charCount := make(map[byte]int)
	for _, row := range matrix {
		for _, ch := range row {
			charCount[ch]++
		}
	}

	for i := 0; i < len(word); i++ {
		if charCount[word[i]] == 0 {
			return false
		}
		charCount[word[i]]--
	}

	return true
}

func mainWS() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println("Does the word exist in the board?", exist(board, word)) // Expected: true

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "SEE"
	fmt.Println("Does the word exist in the board?", exist(board, word)) // Expected: true

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCB"
	fmt.Println("Does the word exist in the board?", exist(board, word)) // Expected: false
}
