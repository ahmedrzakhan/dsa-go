package main

import (
	"fmt"
)

/**

22. Generate Parentheses

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1 <= n <= 8
*/

func generateParentheses(N int) []string {
	var subsets []string
	helperGen(0, 0, N, "", &subsets)
	return subsets
}

func helperGen(open, close, N int, curSet string, subsets *[]string) {
	if len(curSet) == 2*N {
		*subsets = append(*subsets, curSet)
		return
	}

	if open < N {
		helperGen(open+1, close, N, curSet+"(", subsets)
	}
	if close < open {
		helperGen(open, close+1, N, curSet+")", subsets)
	}
}

func mainGen() {
	n := 3
	fmt.Println(generateParentheses(n))
}
