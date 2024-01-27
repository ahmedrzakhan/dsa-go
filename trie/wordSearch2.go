package main

import "fmt"

/**
212. Word Search II

Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent
cells are horizontally or vertically neighboring. The same letter cell may not be used
more than once in a word.

Example 1:

Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
Example 2:

Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []

Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] is a lowercase English letter.
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
All the strings of words are unique.
*/

// Define TrieNode
type TrieNodeWS struct {
	children map[rune]*TrieNodeWS
	word     string
}

// NewTrieNode creates a new Trie node
func NewTrieNodeWS() *TrieNodeWS {
	return &TrieNodeWS{children: make(map[rune]*TrieNodeWS)}
}

// Insert inserts a word into the Trie
func (t *TrieNodeWS) Insert(word string) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = NewTrieNodeWS()
		}
		curr = curr.children[ch]
	}
	curr.word = word // Store the word at the end node
}

// TC - O(M*N*4^L) SC - O(AL)
// findWords finds all words in the given matrix
func findWords(matrix [][]byte, words []string) []string {
	ROWS, COLS := len(matrix), len(matrix[0])
	var result []string

	if ROWS == 0 || COLS == 0 {
		return result
	}

	// Build the Trie
	curr := NewTrieNodeWS()
	for _, word := range words {
		curr.Insert(word)
	}

	// Start DFS from every cell
	for R := range matrix {
		for C := range matrix[R] {
			dfsWS(R, C, ROWS, COLS, curr, &result, matrix)
		}
	}
	return result
}

// dfs performs a Depth-First Search
func dfsWS(R, C, ROWS, COLS int, curr *TrieNodeWS, result *[]string, matrix [][]byte) {
	if R < 0 || R >= ROWS || C < 0 || C >= COLS {
		return // Check bounds
	}

	temp := matrix[R][C]
	// Check if current cell is part of Trie and not visited
	nextNode, exists := curr.children[rune(temp)]
	if !exists || temp == '#' { // '#' is used to mark visited cells
		return
	}

	// Check if a word is found
	if nextNode.word != "" {
		*result = append(*result, nextNode.word)
		nextNode.word = "" // Avoid duplicate entries
	}

	// Mark the current cell as visited
	matrix[R][C] = '#'
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range directions {
		nextRow, nextCol := R+dir[0], C+dir[1]
		dfsWS(nextRow, nextCol, ROWS, COLS, nextNode, result, matrix) // Explore all possible directions
	}
	// Unmark the current cell
	matrix[R][C] = temp
}

func mainWS() {
	matrix := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(matrix, words)) // Expected output: ["eat","oath"] or ["oath","eat"]
}
