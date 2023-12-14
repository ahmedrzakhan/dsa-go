package main

import "fmt"

/**
338. Counting Bits

Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Example 1:

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
Example 2:

Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

Constraints:

0 <= n <= 105

Follow up:

It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in
linear time O(n) and possibly in a single pass?
Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?
*/

/**
NOTE:
Here are the results of n & (n - 1) for numbers from 1 to 16:

1 (001) & 0 (000) = 0
2 (010) & 1 (001) = 0 (power of 2, resets to 0)
3 (011) & 2 (010) = 2
4 (100) & 3 (011) = 0 (power of 2, resets to 0)
5 (101) & 4 (100) = 4
6 (110) & 5 (101) = 4
7 (111) & 6 (110) = 6
8 (1000) & 7 (0111) = 0 (power of 2, resets to 0)
9 (1001) & 8 (1000) = 8
10 (1010) & 9 (1001) = 8
11 (1011) & 10 (1010) = 10
12 (1100) & 11 (1011) = 8
13 (1101) & 12 (1100) = 12
14 (1110) & 13 (1101) = 12
15 (1111) & 14 (1110) = 14
16 (10000) & 15 (01111) = 0 (power of 2, resets to 0)
These results highlight the patterns we discussed earlier. Particularly, every time n is a power of 2,
the result resets to 0. For other numbers, the result is n with its least significant set bit remove
*/

// O(NLogN)
func hamm(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

// TC - O(N), SC - O(N)
// Optimized function using dynamic programming
func hammDP(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1 + dp[i&(i-1)] // Removes the least significant 1-bit from i
	}
	return dp
}

func mainCount() {
	// Brute
	N := 5
	// arr := make([]int, N+1)
	// for i := 0; i <= N; i++ {
	// 	arr[i] = hamm(i)
	// }
	// fmt.Println(arr)
	arr := hammDP(N)
	// arr := countBits(N)
	fmt.Println(arr)
}
