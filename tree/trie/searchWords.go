package main

import "fmt"

/**
211. Design Add and Search Words Data Structure

Design a data structure that supports adding new words and finding if a string matches any
 previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches
word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True

Constraints:

1 <= word.length <= 25
word in addWord consists of lowercase English letters.
word in search consist of '.' or lowercase English letters.
There will be at most 2 dots in word for search queries.
At most 104 calls will be made to addWord and search.
*/

type WordDictionary struct {
	children map[rune]*WordDictionary
	isWord   bool
}

/** Initialize your data structure here. */
func ConstructorSR() WordDictionary {
	return WordDictionary{children: make(map[rune]*WordDictionary)}
}

// TC - O(N), SC - O(N)
/** Adds a word into the data structure. */
func (t *WordDictionary) AddWordSR(word string) {
	curr := t
	for _, ch := range word {
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = &WordDictionary{children: make(map[rune]*WordDictionary)}
		}
		curr = curr.children[ch]
	}
	curr.isWord = true
}

// TC - O(N), SC - O(H)
/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (t *WordDictionary) Search(word string) bool {
	return t.searchHelper(0, word)
}

func (t *WordDictionary) searchHelper(i int, word string) bool {
	if i == len(word) {
		return t.isWord
	}

	ch := rune(word[i])
	if ch == '.' {
		for _, child := range t.children {
			if child.searchHelper(i+1, word) {
				return true
			}
		}
		return false
	} else {
		if child, exists := t.children[ch]; exists {
			return child.searchHelper(i+1, word)
		}
		return false
	}
}

func main() {
	wordDictionary := ConstructorSR()
	// wordDictionary.AddWord("bad")
	// wordDictionary.AddWord("dad")
	// wordDictionary.AddWord("mad")
	// fmt.Println(wordDictionary.Search("pad")) // returns false
	// fmt.Println(wordDictionary.Search("bad")) // returns true
	// fmt.Println(wordDictionary.Search(".ad")) // returns true
	// fmt.Println(wordDictionary.Search("b..")) // returns true

	wordDictionary.AddWordSR("bads")
	wordDictionary.AddWordSR("bamf")
	wordDictionary.AddWordSR("bace")
	fmt.Println(wordDictionary.Search("ba.m")) // returns false
}
