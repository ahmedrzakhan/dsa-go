package main

import "fmt"

/**
https://www.geeksforgeeks.org/find-the-longest-substring-with-k-unique-characters-in-a-given-string/

Given a string you need to print longest possible substring that has exactly M unique characters.
If there is more than one substring of longest possible length, then print any one of them.

Input: Str = “aabbcc”, k = 1
Output: 2
Explanation: Max substring can be any one from {“aa” , “bb” , “cc”}.

Input: Str = “aabbcc”, k = 2
Output: 4
Explanation: Max substring can be any one from {“aabb” , “bbcc”}.

*/

func lengthOfLongestSubstringKDistinct(S string, K int) int {
	if K == 0 {
		return 0
	}

	// Map to store the count of each character in the current window.
	charCountMap := make(map[byte]int)
	left, maxLength := 0, 0

	for right := range S {
		// Increment the count for the current character.
		charCountMap[S[right]]++

		// While we have more than 'k' distinct characters, shrink the window.
		for len(charCountMap) > K {
			// Decrement the count of the leftmost character.
			charCountMap[S[left]]--
			// If count drops to zero, remove the character from the map.
			if charCountMap[S[left]] == 0 {
				delete(charCountMap, S[left])
			}
			left++
		}

		// Update the max length if the current window is larger.
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mainLSK() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2)) // Output: 3
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))    // Output: 2
	fmt.Println(lengthOfLongestSubstringKDistinct("bacc", 2))  // Output: 3
}
