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

// TODO: DP approach
func change2DP(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] += dp[x-coin]
		}
	}

	return dp[amount]
}

// TODO: get TC and SC
// TC - O(N*2^N), SC - O(N * K)
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

func mainChange2() {
	arr := []int{1, 2, 5}
	target := 5
	result := change2(arr, target)
	fmt.Println("Combinations that add up to", target, "are:", result)
}
