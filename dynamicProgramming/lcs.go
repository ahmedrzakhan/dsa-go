package main

import (
	"fmt"
)

/**
1143. Longest Common Subsequence

Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without
changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

Constraints:

1 <= text1.length, text2.length <= 1000
text1 and text2 consist of only lowercase English characters.
*/

// TC - O(R*C), SC - O(R*C)
func longestCommonSubsequence(S1 string, S2 string) int {
	DP := make([][]int, len(S1)+1)

	for R := 0; R < len(DP); R++ {
		DP[R] = make([]int, len(S2)+1)
	}

	for R := len(S1) - 1; R >= 0; R-- {
		for C := len(S2) - 1; C >= 0; C-- {
			if S1[R] == S2[C] {
				DP[R][C] = 1 + DP[R+1][C+1]
			} else {
				DP[R][C] = max(DP[R][C+1], DP[R+1][C])
			}
		}
	}

	return DP[0][0]
}

func mainLc() {
	// Example usage
	S1 := "abcde"
	S2 := "ace"

	lcs := longestCommonSubsequence(S1, S2)

	fmt.Println("Texts:", S1, S2)
	fmt.Println("Longest Common Subsequence:", lcs)
}
