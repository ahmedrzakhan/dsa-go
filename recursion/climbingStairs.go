package main

import "fmt"

/**
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

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

/**
NOTE: the same function can be used for fibonacci as well, just intial values would be 0 and 1
*/

// TC  - O(2^N), SC - O(N)
func climbStairsBr(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairsBr(n-1) + climbStairsBr(n-2)
}

// TC  - O(N), SC - O(1)
func climbStairs(n int) int {
	first, second := 1, 1
	for i := 0; i < n-1; i++ {
		first, second = second, first+second
	}
	return second
}

func mainCl() {
	n := 5 // Example input
	fmt.Println(climbStairs(n))
}
