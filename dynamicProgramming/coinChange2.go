package main

import (
	"fmt"
)

func dfs(i, acc, amount int, coins []int, cache map[[2]int]int) int {
	if acc == amount {
		return 1
	}
	if acc > amount || i == len(coins) {
		return 0
	}
	key := [2]int{i, acc}
	if val, exists := cache[key]; exists {
		return val
	}

	cache[key] = dfs(i, acc+coins[i], amount, coins, cache) + dfs(i+1, acc, amount, coins, cache)
	return cache[key]
}

func changeMemoization(amount int, coins []int) int {
	cache := make(map[[2]int]int)
	return dfs(0, 0, amount, coins, cache)
}

// DP 2d array
func changeDP(amount int, coins []int) int {
	dp := make([][]int, amount+1)
	for i := range dp {
		dp[i] = make([]int, len(coins)+1)
	}

	// Initialize the base case
	for i := range dp[0] {
		dp[0][i] = 1
	}

	// Dynamic programming to fill the dp table
	for a := 1; a <= amount; a++ {
		for i := len(coins) - 1; i >= 0; i-- {
			// Number of combinations without using the current coin
			dp[a][i] = dp[a][i+1]

			// Number of combinations using the current coin
			if a-coins[i] >= 0 {
				dp[a][i] += dp[a-coins[i]][i]
			}
		}
	}

	return dp[amount][0]
}

// DP reduced space
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		nextDP := make([]int, amount+1)
		nextDP[0] = 1

		for a := 1; a <= amount; a++ {
			nextDP[a] = dp[a]
			if a-coin >= 0 {
				nextDP[a] += nextDP[a-coin]
			}
		}
		dp = nextDP
	}

	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println("Number of ways to make change:", changeMemoization(amount, coins))
	fmt.Println("Number of ways to make change:", changeDP(amount, coins))
	fmt.Println("Number of ways to make change:", change(amount, coins))
}
