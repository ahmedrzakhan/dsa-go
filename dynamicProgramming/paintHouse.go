package main

import (
	"fmt"
	"math"
)

/**
515 · Paint House

Description
There are a row of n houses, each house can be painted with one of the three colors: red, blue or green.
The cost of painting each house with a certain color is different. You have to paint all the houses such
that no two adjacent houses have the same color, and you need to cost the least. Return the minimum cost.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example,
costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1
with color green, and so on... Find the minimum cost to paint all houses.

Only $39.9 for the "Twitter Comment System Project Practice" within a limited time of 7 days!

WeChat Notes Twitter for more information（WeChat ID jiuzhangfeifei）

All costs are positive integers.

Example
Example 1:

Input: [[14,2,11],[11,14,5],[14,3,10]]
Output: 10
Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue. Minimum cost: 2 + 5 + 3 = 10.
Example 2:

Input: [[1,2,3],[1,4,6]]
Output: 3
*/

// TC - O(N) , SC - O(1)
func minCost(cost [][]int) int {
	N := len(cost)
	DP := make([][]int, N)
	for i := range DP {
		DP[i] = make([]int, 3)
	}

	// Initialize the base case
	for i := 0; i < 3; i++ {
		DP[0][i] = cost[0][i]
	}

	// Fill out the dp table
	for i := 1; i < N; i++ {
		DP[i][0] = cost[i][0] + min(DP[i-1][1], DP[i-1][2])
		DP[i][1] = cost[i][1] + min(DP[i-1][0], DP[i-1][2])
		DP[i][2] = cost[i][2] + min(DP[i-1][0], DP[i-1][1])
	}

	// Return the minimum cost from the last row
	return min(DP[N-1][0], DP[N-1][1], DP[N-1][2])
}

// TC - O(N * K) , SC - O(1)
func minCostK(cost [][]int) int {
	N := len(cost)
	K := len(cost[0]) // Number of paints

	DP := make([][]int, N)

	for i := range DP {
		DP[i] = make([]int, K)
	}

	// Initialize the base case
	for i := 0; i < K; i++ {
		DP[0][i] = cost[0][i]
	}

	// Fill out the dp table
	for i := 1; i < N; i++ {
		for j := 0; j < K; j++ {
			minPrev := math.MaxInt // Initialize to maximum possible value
			for k := 0; k < K; k++ {
				if k != j && DP[i-1][k] < minPrev {
					minPrev = DP[i-1][k]
				}
			}
			DP[i][j] = cost[i][j] + minPrev
		}
	}

	// Find the minimum cost from the last row
	minCost := math.MaxInt // Initialize to maximum possible value
	for i := 0; i < K; i++ {
		if DP[N-1][i] < minCost {
			minCost = DP[N-1][i]
		}
	}
	return minCost
}

func mainPh() {
	costs := [][]int{{1, 4, 5}, {8, 2, 11}, {3, 5, 1}}
	fmt.Println(minCost(costs)) // Output: 4

	costs = [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println("Minimum cost to paint houses:", minCost(costs)) // Output: 10

	costs = [][]int{{1, 4, 5}, {8, 2, 11}, {3, 5, 1}}
	fmt.Println(minCostK(costs)) // Output: 4

	costs = [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println("Minimum cost to paint houses:", minCostK(costs)) // Output: 10
}
