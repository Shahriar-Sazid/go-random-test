package trie

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/Shahriar-Sazid/go-random-test/ed"
)

type STrieNode struct {
	Char     rune
	IsEnd    bool
	Range    [2]int
	Counter  int
	Children []*STrieNode
}

func (n *STrieNode) findChild(x rune) *STrieNode {
	if n == nil || len(n.Children) == 0 {
		return nil
	}

	s, e := 0, len(n.Children)-1
	m := (s + e) / 2

	for {
		if n.Children[m].Char == x {
			return n.Children[m]
		}
		if s >= m && m >= e {
			return nil
		}

		if x < n.Children[m].Char {
			e = m - 1
			m = (s + e) / 2
		} else if x > n.Children[m].Char {
			s = m + 1
			m = (s + e) / 2
		}
	}
}

func NewSTrieNode(char rune) *STrieNode {
	return &STrieNode{
		Char:     char,
		IsEnd:    false,
		Range:    [2]int{math.MaxInt, -1},
		Counter:  0,
		Children: make([]*STrieNode, 0, 15),
	}
}

type STrie struct {
	Root *STrieNode
}

func NewSTrie() *STrie {
	return &STrie{
		Root: NewSTrieNode(' '),
	}
}

func (t *STrie) Insert(word string) {
	node := t.Root
	for _, letter := range word {
		child := node.findChild(letter)
		if child == nil {
			child = NewSTrieNode(letter)
			node.Children = append(node.Children, child)
			sort.Slice(node.Children, func(i, j int) bool {
				return node.Children[i].Char < node.Children[j].Char
			})
		}
		node.Range = [2]int{min(node.Range[0], len(word)), max(node.Range[1], len(word))}
		node = child
	}
	node.IsEnd = true
	node.Counter++
}

var ranges [][2]int

func init() {
	ranges = [][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {3, 5}, {4, 7}, {5, 8}, {5, 9}, {6, 11}, {7, 12}, {8, 13}, {9, 14}, {9, 15}, {10, 16}, {11, 17}, {12, 19}, {13, 20}}
}
func (t *STrie) FuzzySearch(word string) []FuzzyResult {
	chars := []rune(word)
	firstChar := chars[0]

	start := t.Root.findChild(firstChar)
	if start == nil {
		return nil
	}

	rangesOverlap := func(range1 [2]int, range2 [2]int) bool {
		start1, end1 := range1[0], range1[1]
		start2, end2 := range2[0], range2[1]
		return start1 <= end2 && end1 >= start2
	}

	getRange := func(length int) [2]int {
		if length >= len(ranges) {
			return [2]int{13, 100}
		}
		return ranges[length]
	}

	results := make([]FuzzyResult, 0, 100)
	var fuzzyDFS func(*STrieNode, int, string)
	fuzzyDFS = func(node *STrieNode, level int, pathVisited string) {
		if node.IsEnd {
			matchRatio := ed.IncrementalMatchRatio(pathVisited, word, level, func() int {
				if len(pathVisited) >= len(word) {
					return 0
				}
				return len(word) - len(pathVisited)
			}())
			if (len(word) >= 0 && len(word) < 4 && matchRatio >= 0.9) ||
				(len(word) >= 4 && len(word) < 6 && matchRatio >= 0.80) ||
				(len(word) >= 6 && len(word) < 8 && matchRatio >= 0.77) ||
				(len(word) >= 8 && matchRatio >= 0.75) {
				{
					results = append(results, FuzzyResult{
						Word:  word,
						Token: pathVisited,
						Ratio: matchRatio,
					})
				}
			}
		}

		for _, nextNode := range node.Children {
			incrementalMR := ed.IncrementalMatchRatio(pathVisited+string(nextNode.Char), word, level, 1)
			if ((len(pathVisited) >= 0 && len(pathVisited) < 4 && incrementalMR >= 0.25) ||
				(len(pathVisited) >= 4 && len(pathVisited) < 6 && incrementalMR >= 0.4) ||
				(len(pathVisited) >= 6 && len(pathVisited) < 8 && incrementalMR >= 0.5) ||
				(len(pathVisited) >= 8 && incrementalMR >= 0.55)) &&
				rangesOverlap(node.Range, getRange(len(word))) {
				fuzzyDFS(nextNode, level+1, pathVisited+string(nextNode.Char))
			}
		}
	}

	ed.IncrementalMatchRatio(string(firstChar), string(firstChar), 0, 1)
	fuzzyDFS(start, 1, string(firstChar))
	return results
}

func TestTrieFuzz() {
	file, err := os.Open("place_mapping.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	words := make([]string, 0, 30000)
	exceptAlphabet, err := regexp.Compile("[^a-zA-Z]")
	if err != nil {
		return
	}
	for _, record := range records {
		districts := string(exceptAlphabet.ReplaceAll([]byte(record[2]), []byte(" ")))
		zones := string(exceptAlphabet.ReplaceAll([]byte(record[8]), []byte(" ")))
		areas := string(exceptAlphabet.ReplaceAll([]byte(record[12]), []byte(" ")))
		words = append(words, strings.Fields(strings.ToLower(districts))...)
		words = append(words, strings.Fields(strings.ToLower(zones))...)
		words = append(words, strings.Fields(strings.ToLower(areas))...)
	}
	t1 := NewSTrie()
	t2 := NewTrie()

	for _, word := range words {
		t1.Insert(word)
		t2.Insert(word)
	}

	word := "siddirganj"
	startTime := time.Now()
	results := t1.FuzzySearch(word)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Microseconds()
	fmt.Printf("strie fuzz took %d us to search %s\n", elapsedTime, word)
	for _, result := range results {
		fmt.Println(result.Word, result.Token, result.Ratio)
	}

	startTime = time.Now()
	results = t2.FuzzySearch(word)
	endTime = time.Now()
	elapsedTime = endTime.Sub(startTime).Microseconds()
	fmt.Printf("trie fuzz took %d us to search %s\n", elapsedTime, word)
	for _, result := range results {
		fmt.Println(result.Word, result.Token, result.Ratio)
	}
}
