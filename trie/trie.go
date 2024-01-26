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

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = NewTrieNode()
		}
		curr = curr.children[c]
	}
	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return curr.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, c := range prefix {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
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
