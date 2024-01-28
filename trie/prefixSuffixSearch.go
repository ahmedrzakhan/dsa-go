package main

import "fmt"

/**
745. Prefix and Suffix Search

Design a special dictionary that searches the words in it by a prefix and a suffix.

Implement the WordFilter class:

WordFilter(string[] words) Initializes the object with the words in the dictionary.
f(string pref, string suff) Returns the index of the word in the dictionary, which has
the prefix pref and the suffix suff. If there is more than one valid index, return the
largest of them. If there is no such word in the dictionary, return -1.


Example 1:

Input
["WordFilter", "f"]
[[["apple"]], ["a", "e"]]
Output
[null, 0]
Explanation
WordFilter wordFilter = new WordFilter(["apple"]);
wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a"
and suffix = "e".


Constraints:

1 <= words.length <= 104
1 <= words[i].length <= 7
1 <= pref.length, suff.length <= 7
words[i], pref and suff consist of lowercase English letters only.
At most 104 calls will be made to the function f.
*/

type TrieNodePS struct {
	children map[string]*TrieNodePS
	weight   int
}

func NewTrieNodePS() *TrieNodePS {
	return &TrieNodePS{children: make(map[string]*TrieNodePS)}
}

type WordFilter struct {
	root *TrieNodePS
}

// TC - O(N*K) SC - O(N*K)
func ConstructorPS(words []string) WordFilter {
	wf := WordFilter{root: NewTrieNodePS()}
	for weight, word := range words {
		key := "{" + word + "}" // The key is the whole word prefixed by '{'

		// Insert the whole word with '{' prefix once
		wf.root.Insert(key, weight)

		// Iterate over the word to insert each suffix combined with the '{word}'
		for j := len(word) - 1; j >= 0; j-- {
			suffixKey := word[j:] + key // Suffix of the word combined with '{word}'
			wf.root.Insert(suffixKey, weight)
		}
	}
	return wf
}

func (t *TrieNodePS) Insert(word string, weight int) {
	curr := t
	for _, char := range word {
		ch := string(char)
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = NewTrieNodePS()
		}
		curr = curr.children[ch]
		curr.weight = weight
	}
}

func (t *WordFilter) F(prefix string, suffix string) int {
	curr := t.root
	for _, ch := range suffix + "{" + prefix {
		c := string(ch)
		if _, exists := curr.children[c]; !exists {
			return -1
		}
		curr = curr.children[c]
	}
	return curr.weight
}

func mainPS() {
	words := []string{"apple", "applet"}
	wf := ConstructorPS(words)
	fmt.Println(wf.F("a", "e")) // Output: 0
}
