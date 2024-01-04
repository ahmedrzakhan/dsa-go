package main

import "fmt"

/**
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen
and an empty space, respectively.

Example 1:

Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
Example 2:

Input: n = 1
Output: [["Q"]]

Constraints:

1 <= n <= 9
*/

// TC - O(N!), SC - O(S(N)*N^2)
func solveNQueens(boardSize int) [][]string {
	var solutions [][]string
	var currentBoard []string
	columnHits := make(map[int]int)
	posDiag := make(map[int]int)
	negDiag := make(map[int]int)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == boardSize {
			solutions = append(solutions, append([]string{}, currentBoard...))
			return
		}

		for col := 0; col < boardSize; col++ {
			if columnHits[col] > 0 || posDiag[col+row] > 0 || negDiag[col-row] > 0 {
				continue
			}
			columnHits[col], posDiag[col+row], negDiag[col-row] = 1, 1, 1
			rowString := make([]byte, boardSize)
			for i := range rowString {
				if i == col {
					rowString[i] = 'Q'
				} else {
					rowString[i] = '.'
				}
			}
			currentBoard = append(currentBoard, string(rowString))
			backtrack(row + 1)
			columnHits[col], posDiag[col+row], negDiag[col-row] = 0, 0, 0
			currentBoard = currentBoard[:len(currentBoard)-1]
		}
	}

	backtrack(0)
	return solutions
}

func mainNQueens() {
	boardSize := 4
	solutions := solveNQueens(boardSize)
	for _, solution := range solutions {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
