package main

import (
	"fmt"
)

/**
518. Coin Change II

You are given an integer array coins representing coins of different denominations and an integer
amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be
made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.

The answer is guaranteed to fit into a signed 32-bit integer.

Example 1:

Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:

Input: amount = 10, coins = [10]
Output: 1

Constraints:

1 <= coins.length <= 300
1 <= coins[i] <= 5000
All the values of coins are unique.
0 <= amount <= 5000
*/

func change2(arr []int, target int) int {
	count := 0
	if len(arr) == 0 {
		return count
	}
	curSet := make([]int, 0)
	helperChange2(0, arr, target, 0, &curSet, &count)
	return count
}

// TC - O(2^N), SC - O(N)
func helperChange2(idx int, arr []int, target int, curSum int, curSet *[]int, count *int) {
	if curSum == target {
		*count++
		return
	}
	if curSum > target {
		return
	}
	for i := idx; i < len(arr); i++ {
		*curSet = append(*curSet, arr[i])
		helperChange2(i, arr, target, curSum+arr[i], curSet, count)
		*curSet = (*curSet)[:len(*curSet)-1]
	}
}

// memoization
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
	fmt.Println("Number of ways to make change:", change2(coins, amount))
}
