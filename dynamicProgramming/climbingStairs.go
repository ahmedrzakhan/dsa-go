package main

import (
	"fmt"
)

/**
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct
ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:

1 <= n <= 45
*/

// TC - O(2^N), SC - O(N)
// Brute Force
func climbStairsBruteForce(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairsBruteForce(n-1) + climbStairsBruteForce(n-2)
}

// TC - O(N), SC - O(N)
// Memoization
func climbStairsMemoization(n int, cache map[int]int) int {
	if n <= 2 {
		return n
	}
	if val, ok := cache[n]; ok {
		return val
	}
	cache[n] = climbStairsMemoization(n-1, cache) + climbStairsMemoization(n-2, cache)
	return cache[n]
}

// TC - O(N), SC - O(1)
// Dynamic Programming
func climbStairsDP(n int) int {
	if n <= 2 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

func mainCSt() {
	n := 5 // Example input
	fmt.Printf("Brute Force: %d\n", climbStairsBruteForce(n))

	memo := make(map[int]int)
	fmt.Printf("Memoization: %d\n", climbStairsMemoization(n, memo))

	fmt.Printf("Dynamic Programming: %d\n", climbStairsDP(n))
}
