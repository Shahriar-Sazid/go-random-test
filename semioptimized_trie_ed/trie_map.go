package semioptimizedtrieed

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	Char     rune
	IsEnd    bool
	Counter  int
	Children map[rune]*TrieNode
}

func NewTrieNode(char rune) *TrieNode {
	return &TrieNode{
		Char:     char,
		IsEnd:    false,
		Counter:  0,
		Children: make(map[rune]*TrieNode),
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(' '),
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, syllable := range word {
		child, ok := node.Children[syllable]
		if !ok {
			child = NewTrieNode(syllable)
			node.Children[syllable] = child
		}
		node = child
	}
	node.IsEnd = true
	node.Counter++
}

func (t *Trie) DFS(node *TrieNode, prefix string, output *[]pair) {
	if node.IsEnd {
		*output = append(*output, pair{prefix + string(node.Char), node.Counter})
	}

	for _, child := range node.Children {
		t.DFS(child, prefix+string(node.Char), output)
	}
}

func (t *Trie) FuzzySearch(word string) []FuzzyResult {
	chars := []rune(word)
	firstChar := chars[0]

	start := t.Root.Children[firstChar]
	if start == nil {
		return nil
	}

	results := make([]FuzzyResult, 0, 100)
	var fuzzyDFS func(*TrieNode, int, string)
	fuzzyDFS = func(node *TrieNode, level int, pathVisited string) {
		if node.IsEnd {
			matchRatio := IncrementalMatchRatio2D(pathVisited, word, level, func() int {
				if len(pathVisited) >= len(word) {
					return 0
				}
				return len(word) - len(pathVisited)
			}())
			if matchRatio >= 0.75 {
				results = append(results, FuzzyResult{
					Word:  word,
					Token: pathVisited,
					Ratio: matchRatio,
				})
			}
		}

		for nextChar, nextNode := range node.Children {
			if IncrementalMatchRatio2D(pathVisited+string(nextChar), word, level, 1) >= 0.3 {
				fuzzyDFS(nextNode, level+1, pathVisited+string(nextChar))
			}
		}
	}

	IncrementalMatchRatio2D(string(firstChar), string(firstChar), 0, 1)
	fuzzyDFS(start, 1, string(firstChar))

	return results
}

func (t *Trie) Query(x string) []pair {
	output := make([]pair, 0)
	node := t.Root

	for _, char := range x {
		child, ok := node.Children[char]
		if !ok {
			return []pair{}
		}
		node = child
	}

	t.DFS(node, x[:len(x)-1], &output)

	sort.Slice(output, func(i, j int) bool {
		return output[i].Count > output[j].Count
	})

	return output
}

type pair struct {
	Word  string
	Count int
}

func TestTrie() {
	trie := NewTrie()
	words := []string{"apple", "app", "apricot", "banana", "bat", "batman"}

	for _, word := range words {
		trie.Insert(word)
	}

	results := trie.Query("ap")
	for _, result := range results {
		fmt.Println(result.Word, result.Count)
	}
}
