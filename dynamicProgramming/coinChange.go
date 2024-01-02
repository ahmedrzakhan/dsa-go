package main

import "fmt"

/**
322. Coin Change

You are given an integer array coins representing coins of different denominations and an integer amount
representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be
made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0

Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
*/

/**
NOTE:
[1, 3, 4, 5] amount = 7 (is it really greedy, start with largest amount?)
 5 + 1 + 1 is 3 coins
 4 + 3 is 2 coins
So, in summary, this block of code is filling up the changes array, where each entry changes[i] will
ultimately hold the minimum number of coins needed to make up the amount i.
It does this by considering each coin denomination for each amount and updating the minimum
number of coins needed. This is a bottom-up dynamic programming approach, where solutions to smaller
sub-problems (minimum coins for smaller amounts) are used to solve larger problems (minimum coins
for the total amount).
ASK: interviewer if number of coins are infinite
*/

// TC - O(amount * N), SC - O(amount)
func coinChange(coins []int, amount int) int {
	changes := make([]int, amount+1)

	for i := 1; i < len(changes); i++ {
		changes[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				changes[i] = min(changes[i], 1+changes[i-coin])
			}
		}
	}
	if changes[amount] != amount+1 {
		return changes[amount]
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func mainCoin() {
	// coins := []int{1, 2, 5}
	coins := []int{2}
	amount := 3
	fmt.Printf("Minimum coins needed: %d\n", coinChange(coins, amount))
}
