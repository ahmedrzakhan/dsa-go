package main

import "fmt"

/**
714. Best Time to Buy and Sell Stock with Transaction Fee

You are given an array prices where prices[i] is the price of a given stock on the ith day,
and an integer fee representing a transaction fee.

Find the maximum profit you can achieve. You may complete as many transactions as you like,
but you need to pay the transaction fee for each transaction.

Note:

You may not engage in multiple transactions simultaneously (i.e., you must sell the stock
before you buy again).
The transaction fee is only charged once for each stock purchase and sale.

Example 1:

Input: prices = [1,3,2,8,4,9], fee = 2
Output: 8
Explanation: The maximum profit can be achieved by:
- Buying at prices[0] = 1
- Selling at prices[3] = 8
- Buying at prices[4] = 4
- Selling at prices[5] = 9
The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
Example 2:

Input: prices = [1,3,7,5,10,3], fee = 3
Output: 6


Constraints:

1 <= prices.length <= 5 * 104
1 <= prices[i] < 5 * 104
0 <= fee < 5 * 104
*/

// TC - O(N), SC - O(1)
func maxProfit(A []int, fee int) int {
	N := len(A)
	cash := 0     // Maximum profit if we end the day without holding a stock
	hold := -A[0] // Maximum profit if we end the day holding a stock

	for i := 1; i < N; i++ {
		cash = max(cash, hold+A[i]-fee) // Sell the stock (previous holding state)
		hold = max(hold, cash-A[i])     // Buy a stock (previous cash state)
	}

	return cash // Final result is the maximum profit without holding a stock
}

func mainMX() {
	A := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	maxProfit := maxProfit(A, fee)
	fmt.Println("Maximum Profit with Transaction Fee:", maxProfit)
}
