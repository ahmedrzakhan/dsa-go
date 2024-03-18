package main

import (
	"fmt"
	"strings"
)

/**
49. Group Anagrams

Given an array of strings strs, group the anagrams together.
You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters
of a different word or phrase, typically using all the original
letters exactly once.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]

Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

/**
brute force
sort each key keep it in a hashmap, if key matches again
push it in same hash array
{
sortedKey: [ item1, item2 ]
sortedKey2: [ item3, item4 ]
}

TC - O(N * K LogK), SC - O(NK)
*/

// TC - O(N * K), SC - O(NK)
func groupAnagrams(A []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, str := range A {
		var count [26]int

		for _, ch := range str {
			lCh := strings.ToLower(string(ch))[0]
			count[lCh-'a']++
		}

		anagramMap[count] = append(anagramMap[count], str)
	}

	result := make([][]string, len(anagramMap))
	idx := 0

	for _, group := range anagramMap {
		result[idx] = group
		idx++
	}

	return result
}

func mainGA() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupedAnagrams := groupAnagrams(strs)

	fmt.Println("Grouped Anagrams:", groupedAnagrams)

	strs = []string{"Eat", "tea", "tan", "ate", "nat", "bat"}
	groupedAnagrams = groupAnagrams(strs)

	fmt.Println("Grouped Anagrams:", groupedAnagrams)
}
