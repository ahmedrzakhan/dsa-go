package main

import "fmt"

/**
77. Combinations

Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.

Example 1:

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
Example 2:

Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.

Constraints:

1 <= n <= 20
1 <= k <= n
*/

// TC: O(k * C(n, k))
func combinations(n, k int) [][]int {
	var combs [][]int
	helperC(1, []int{}, &combs, n, k)
	return combs
}

// Given n=5 and  k=2, we get 5 ! 2 ! ( 5 − 2 ) ! 2!(5−2)! 5! ​ = 10 10

// Helper function for recursion
func helperC(i int, curComb []int, combs *[][]int, n, k int) {
	if len(curComb) == k {
		*combs = append(*combs, append([]int(nil), curComb...))
		return
	}
	if i > n {
		return
	}

	// Decision to include i
	helperC(i+1, append(curComb, i), combs, n, k)

	// Decision NOT to include i
	helperC(i+1, curComb, combs, n, k)
}

func mainC() {
	n := 4
	k := 2
	result := combinations(n, k)
	fmt.Println("Combinations are:", result)
}
