package main

import "fmt"

// Problem Statement:
// Given a list of words, implement a Trie and provide methods to search for a word, search for a prefix,
// and count the number of words with a given prefix.

type TrieNode struct {
	Children  map[rune]*TrieNode
	IsEnd     bool
	WordCount int
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, c := range word {
		if node.Children[c] == nil {
			node.Children[c] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}

		node = node.Children[c]
		node.WordCount++
	}

	node.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		if node.Children[c] == nil {
			return false
		}

		node = node.Children[c]
	}

	return node.IsEnd
}

func (t *Trie) StartsWith(word string) bool {
	node := t.root
	for _, c := range word {
		if node.Children[c] == nil {
			return false
		}

		node = node.Children[c]
	}

	return true
}

func (t *Trie) CountStartsWith(word string) int {
	node := t.root
	for _, char := range word {
		if node.Children[char] == nil {
			return 0
		}
		node = node.Children[char]
	}
	return node.WordCount
}

func main() {
	trie := NewTrie()

	words := []string{"apple", "app", "apricot", "banana", "bat"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println("Search 'apple':", trie.Search("apple"))     // true
	fmt.Println("Search 'app':", trie.Search("app"))         // true
	fmt.Println("Search 'apricot':", trie.Search("apricot")) // true
	fmt.Println("Search 'orange':", trie.Search("orange"))   // false

	fmt.Println("StartsWith 'ap':", trie.StartsWith("ap"))   // true
	fmt.Println("StartsWith 'ora':", trie.StartsWith("ora")) // false

	fmt.Println("CountWordsWithPrefix 'ap':", trie.CountStartsWith("ap"))   // 3
	fmt.Println("CountWordsWithPrefix 'ban':", trie.CountStartsWith("ban")) // 1
}
