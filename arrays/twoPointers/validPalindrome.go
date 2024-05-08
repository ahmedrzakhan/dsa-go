package main

import (
	"strings"
	"unicode"
)

/**
125. Valid Palindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

func isPalindrome(s string) bool {
	var cleanedStr strings.Builder
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			cleanedStr.WriteRune(unicode.ToLower(ch))
		}
	}

	cleanedString := cleanedStr.String() // Get the accumulated string
	lengthOfStr := len(cleanedString)    // Get the length of the string

	for i := 0; i < lengthOfStr/2; i++ {
		// NOTE: len() - 1- i for checking palindrome
		if cleanedString[i] != cleanedString[lengthOfStr-1-i] {
			return false
		}
	}

	return true
}
