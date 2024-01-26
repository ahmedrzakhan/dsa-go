package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// TC - O(N),SC - O(N)
func (t *Trie) ImInsert(word string) {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = NewTrieNode()
		}
		curr = curr.children[ch]
	}
	curr.isWord = true
}

// TC - O(N),SC - O(1)
func (t *Trie) ImSearch(word string) bool {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isWord
}

// TC - O(N),SC - O(1)
func (t *Trie) ImStartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}

func mainTrie() {
	trie := NewTrie()
	trie.Insert("hello")
	fmt.Println(trie.Search("hello"))      // returns true
	fmt.Println(trie.StartsWith("hell"))   // returns true
	fmt.Println(trie.Search("hell"))       // returns false
	fmt.Println(trie.StartsWith("helloo")) // returns false
	fmt.Println(trie.Search("hell"))       // returns false
}
