package main

import "fmt"

func helperB(R, C int, count *int, ROWS, COLS int) {
	if R == ROWS-1 && C == COLS-1 {
		*count++
		return
	}

	if C < COLS-1 {
		helperB(R, C+1, count, ROWS, COLS)
	}

	if R < ROWS-1 {
		helperB(R+1, C, count, ROWS, COLS)
	}
}

// TC - O(R*C), SC - O(R*C)
func helperM(R, C int, cache *[][]int, ROWS, COLS int) int {
	// Base case: Destination reached
	if R == ROWS-1 && C == COLS-1 {
		return 1
	}

	// If already calculated, return the cached value
	if (*cache)[R][C] != 0 {
		return (*cache)[R][C]
	}

	// Initialize paths count
	paths := 0

	// Explore rightward move if within bounds
	if C < COLS-1 {
		paths += helperM(R, C+1, cache, ROWS, COLS)
	}

	// Explore downward move if within bounds
	if R < ROWS-1 {
		paths += helperM(R+1, C, cache, ROWS, COLS)
	}

	// Update cache with the sum of paths from downward and rightward moves
	(*cache)[R][C] = paths

	return paths
}

func helperD(ROWS, COLS int) int {
	prevRow := make([]int, COLS)

	for R := ROWS - 1; R >= 0; R-- {
		curRow := make([]int, COLS)
		curRow[COLS-1] = 1 // Base case: Rightmost cell has 1 path

		for C := COLS - 2; C >= 0; C-- {
			curRow[C] = curRow[C+1] + prevRow[C]
		}

		prevRow = curRow
	}
	return prevRow[0] // Top-left corner contains the total paths
}

func maindp() {
	count := 0
	// helperB(0, 0, &count, 4, 4)
	fmt.Println(count)

	cache := make([][]int, 4)
	for i := range cache {
		cache[i] = make([]int, 4)
	}

	// fmt.Println(helperM(0, 0, &cache, 4, 4))
	fmt.Println(helperD(4, 4))
}
