package main

import (
	"fmt"
)

/**
17. Letter Combinations of a Phone Number

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number
could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map
to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]

Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/

// TC - O(3^N * 4^M), SC - O(3^N * 4^M)
func letterCombinations(digits string) []string {
	subsets := make([]string, 0)
	if len(digits) == 0 {
		return subsets
	}

	digitToCharMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	curSet := make([]byte, 0)
	helperLC(0, digits, digitToCharMap, &curSet, &subsets)
	return subsets
}

func helperLC(idx int, digits string, digitToCharMap map[byte]string, curSet *[]byte, subsets *[]string) {
	if len(*curSet) == len(digits) {
		*subsets = append(*subsets, string(*curSet))
		return
	}
	digit := digits[idx]
	str := digitToCharMap[digit]
	for i := 0; i < len(str); i++ {
		*curSet = append(*curSet, str[i])
		helperLC(idx+1, digits, digitToCharMap, curSet, subsets)
		*curSet = (*curSet)[:len(*curSet)-1]
	}
}

func maiLetter() {
	digits := "234"
	combinations := letterCombinations(digits)
	fmt.Println("Combinations for", digits, "are:", combinations)
}
