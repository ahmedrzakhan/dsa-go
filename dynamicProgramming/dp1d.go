package main

import (
	"fmt"
)

// Brute Force
func bruteForce(n int) int {
	if n <= 1 {
		return n
	}
	return bruteForce(n-1) + bruteForce(n-2)
}

// Memoization
func memoization(n int, cache map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = memoization(n-1, cache) + memoization(n-2, cache)
	return cache[n]
}

// Dynamic Programming
func dp(n int) int {
	if n < 2 {
		return n
	}

	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}
	return dp[1]
}

func mainDp1d() {
	n := 10
	fmt.Printf("Brute Force: %d\n", bruteForce(n))

	cache := make(map[int]int)
	fmt.Printf("Memoization: %d\n", memoization(n, cache))

	fmt.Printf("Dynamic Programming: %d\n", dp(n))
}
