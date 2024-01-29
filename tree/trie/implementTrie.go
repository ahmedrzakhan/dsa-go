package main

import "fmt"

/**
208. Implement Trie (Prefix Tree)

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store
and retrieve keys in a dataset of strings. There are various applications of this data structure,
 such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted
	 before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word
that has the prefix prefix, and false otherwise.

Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

Constraints:

1 <= word.length, prefix.length <= 2000
word and prefix consist only of lowercase English letters.
At most 3 * 104 calls in total will be made to insert, search, and startsWith.
*/

type TrieNodeIm struct {
	children map[rune]*TrieNode
	isWord   bool
}

type TrieIm struct {
	root *TrieNode
}

func NewTrieIm() *Trie {
	return &Trie{root: NewTrieNode()}
}

func NewTrieNodeIm() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode), isWord: false}
}

// isWord is assgined as false by default in golang

func Constructor() Trie {
	return Trie{
		root: NewTrieNode(),
	}
}

// TC - O(N),SC - O(N)
func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = NewTrieNode()
		}
		curr = curr.children[ch]
	}
	curr.isWord = true
}

// TC - O(N),SC - O(1)
func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, ch := range word {
		if _, exists := curr.children[ch]; !exists {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isWord
}

// TC - O(N),SC - O(1)
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		if _, exists := curr.children[ch]; !exists {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}

func mainImT() {
	trie := NewTrie()
	trie.Insert("hello")
	fmt.Println(trie.Search("hello"))      // returns true
	fmt.Println(trie.StartsWith("hell"))   // returns true
	fmt.Println(trie.Search("hell"))       // returns false
	fmt.Println(trie.StartsWith("helloo")) // returns false
	fmt.Println(trie.Search("hell"))       // returns false
}
