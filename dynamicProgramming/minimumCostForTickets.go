package main

import (
	"fmt"
)

/**
983. Minimum Cost For Tickets

You have planned some train traveling one year in advance. The days of the year in which you will
travel are given as an integer array days. Each day is an integer from 1 to 365.

Train tickets are sold in three different ways:

a 1-day pass is sold for costs[0] dollars,
a 7-day pass is sold for costs[1] dollars, and
a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.

For example, if we get a 7-day pass on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.
Return the minimum number of dollars you need to travel every day in the given list of days.

Example 1:

Input: days = [1,4,6,7,8,20], costs = [2,7,15]
Output: 11
Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
In total, you spent $11 and covered all the days of your travel.
Example 2:

Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
Output: 17
Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
In total, you spent $17 and covered all the days of your travel.

Constraints:

1 <= days.length <= 365
1 <= days[i] <= 365
days is in strictly increasing order.
costs.length == 3
1 <= costs[i] <= 1000
*/

func mincostTicketsBr(A []int, costs []int, dayIdx int) int {
	if dayIdx >= len(A) {
		return 0
	}

	// Try a 1-day pass
	option1 := costs[0] + mincostTicketsBr(A, costs, dayIdx+1)

	// Try a 7-day pass
	i := dayIdx
	for i < len(A) && A[i] < A[dayIdx]+7 {
		i += 1
	}
	option2 := costs[1] + mincostTicketsBr(A, costs, i)

	// Try a 30-day pass
	i = dayIdx
	for i < len(A) && A[i] < A[dayIdx]+30 {
		i += 1
	}
	option3 := costs[2] + mincostTicketsBr(A, costs, i)

	return min(option1, option2, option3)
}

func mincostTicketsM(A []int, costs []int) int {
	memo := make([]int, len(A))
	for i := range memo {
		memo[i] = -1
	}
	return mincostTicketsMemo(A, costs, 0, memo)
}

func mincostTicketsMemo(A []int, costs []int, dayIdx int, memo []int) int {
	if dayIdx >= len(A) {
		return 0
	}

	if memo[dayIdx] != -1 {
		return memo[dayIdx]
	}

	// Try a 1-day pass
	option1 := costs[0] + mincostTicketsMemo(A, costs, dayIdx+1, memo)

	// Try a 7-day pass
	i := dayIdx
	for i < len(A) && A[i] < A[dayIdx]+7 {
		i += 1
	}
	option2 := costs[1] + mincostTicketsMemo(A, costs, i, memo)

	// Try a 30-day pass
	i = dayIdx
	for i < len(A) && A[i] < A[dayIdx]+30 {
		i += 1
	}
	option3 := costs[2] + mincostTicketsMemo(A, costs, i, memo)

	memo[dayIdx] = min(option1, option2, option3)
	return memo[dayIdx]
}

// TC - O(N), SC - O(N)
func mincostTickets(A []int, costs []int) int {
	lastDay := A[len(A)-1] // Find the maximum travel day
	DP := make([]int, lastDay+1)

	travelDays := map[int]bool{}
	for _, day := range A { // Mark all the travel days
		travelDays[day] = true
	}

	for i := 1; i <= lastDay; i++ {
		if !travelDays[i] { // Optimization: Skip non-travel days
			DP[i] = DP[i-1]
			continue
		}

		// 1-day pass
		oneDay := DP[max(0, i-1)] + costs[0]

		// 7-day pass
		sevenDays := DP[max(0, i-7)] + costs[1]

		// 30-day pass
		thirtyDays := DP[max(0, i-30)] + costs[2]

		DP[i] = min(oneDay, sevenDays, thirtyDays) // Pick the minimal cost
	}

	return DP[lastDay]
}

func mainMC() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	minCost := mincostTickets(days, costs)
	fmt.Println("Minimum Cost:", minCost)

	minCost = mincostTicketsBr(days, costs, 0)
	fmt.Println("Minimum Cost:", minCost)
}
