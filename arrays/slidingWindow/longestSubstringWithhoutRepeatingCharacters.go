package main

import "fmt"

/**
3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

// TC - O(N), SC - O(N)
func lengthOfLongestSubstring(S string) int {
	charIndexMap := make(map[byte]int)
	maxLength, L, N := 0, 0, len(S)
	// Using 'L' for the start of the window
	// Using 'R' for the end of the window

	for R := 0; R < N; R++ {
		if prevI, exists := charIndexMap[S[R]]; exists {
			if prevI >= L {
				L = prevI + 1 // if character repeats again move window staring point
			}
		}

		charIndexMap[S[R]] = R

		if maxLength < R-L+1 {
			maxLength = R - L + 1
		}
	}

	return maxLength
}

func mainLSR() {
	inputString := "abcabcbb"
	fmt.Println("Length of the longest substring without repeating characters:", lengthOfLongestSubstring(inputString))
}
