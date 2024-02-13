package ed

import (
	"fmt"
)

var memoArray [100][100]float32

var replaceCost [26][26]float32
var insertionCost [26][26]float32
var deletionCost [26]float32

func getHomophoneCost() [26][26]float32 {
	var homophoneCost [26][26]float32
	for i := 0; i < len(homophoneCost); i++ {
		for j := 0; j < len(homophoneCost[i]); j++ {
			homophoneCost[i][j] = 1
		}
	}

	type homophoneGroup struct {
		group string
		cost  float32
	}
	homophoneGroups := []homophoneGroup{
		{"jgz", 0.5},
		{"ao", 0.5},
		{"ay", 0.5},
		{"ey", 0.6},
		{"ae", 0.5},
		{"ow", 0.5},
		{"uw", 0.5},
		{"uo", 0.5},
		{"ie", 0.5},
		{"yi", 0.5},
		{"ckq", 0.5},
		{"fp", 0.5},
		{"vb", 0.5},
	}

	for _, group := range homophoneGroups {
		for _, c1 := range group.group {
			for _, c2 := range group.group {
				if c1 == c2 {
					continue
				}
				c1i := c1 - 'a'
				c2i := c2 - 'a'
				homophoneCost[c1i][c2i] = group.cost
			}
		}
	}
	return homophoneCost
}

func getAdjacentKeyCost() [26][26]float32 {
	var adjacentKeyCost [26][26]float32
	for i := 0; i < len(adjacentKeyCost); i++ {
		for j := 0; j < len(adjacentKeyCost[i]); j++ {
			adjacentKeyCost[i][j] = 1
		}
	}
	isAlpha := func(s string) bool {
		for _, c := range s {
			if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
				return false
			}
		}
		return true
	}
	qwertyKeyboard := [][]string{
		{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "[", "]", "\\"},
		{"a", "s", "d", "f", "g", "h", "j", "k", "l", ";", "'"},
		{"z", "x", "c", "v", "b", "n", "m", ",", ".", "/"},
	}

	for i := 0; i < len(qwertyKeyboard); i++ {
		for j := 0; j < len(qwertyKeyboard[i]); j++ {
			letter := qwertyKeyboard[i][j]
			var prev string
			if j >= 1 {
				prev = qwertyKeyboard[i][j-1]
			}

			if prev != "" && isAlpha(prev) && isAlpha(letter) {
				p := []rune(prev)[0] - 'a'
				l := []rune(letter)[0] - 'a'
				adjacentKeyCost[p][l] = 0.7
			}
		}
	}
	return adjacentKeyCost
}

func init() {
	for i := 0; i < len(memoArray[0]); i++ {
		memoArray[0][i] = float32(i)
	}
	for i := 0; i < len(memoArray); i++ {
		memoArray[i][0] = float32(i)
	}

	homophoneCost := getHomophoneCost()
	adjacentKeyCost := getAdjacentKeyCost()

	for i := 0; i < len(replaceCost); i++ {
		for j := 0; j < len(replaceCost[i]); j++ {
			replaceCost[i][j] = adjacentKeyCost[i][j]
			replaceCost[i][j] = min(adjacentKeyCost[i][j], homophoneCost[i][j])
		}
	}

	for i := 0; i < len(insertionCost); i++ {
		for j := 0; j < len(insertionCost[i]); j++ {
			insertionCost[i][j] = adjacentKeyCost[i][j]
		}
	}
	insertionCost['s'-'a']['h'-'a'] = 0.3
	insertionCost['t'-'a']['h'-'a'] = 0.3
	insertionCost['g'-'a']['h'-'a'] = 0.3
	insertionCost['b'-'a']['h'-'a'] = 0.3

	deletionCost['h'-'a'] = 0.5
	fmt.Println(deletionCost)
}

func IncrementalED(s, t string, progressSoFar, steps int) float32 {
	distance := memoArray[len(t)][len(s)]
	runeS, runeT := []rune(s), []rune(t)
	for i := progressSoFar; i < progressSoFar+steps; i++ {
		distance = incrementalED(runeS[:min(len(runeS), i+1)], runeT[:min(len(runeT), i+1)])
	}
	return distance
}

func IncrementalMatchRatio(s, t string, progressSoFar, steps int) float32 {
	distance := IncrementalED(s, t, progressSoFar, steps)
	maxLen := progressSoFar + steps

	ratio := 1 - (distance / float32(maxLen))
	// fmt.Printf("%d. match ratio of %s,%s is %f\n", counter, s[:min(len(s), progressSoFar+steps)], t[:min(len(t), progressSoFar+steps)], ratio)
	return ratio
}

func incrementalED(s, t []rune) float32 {
	if len(s) >= len(t) {
		for i := 1; i < len(t); i++ {
			if s[len(s)-1] == t[i-1] {
				memoArray[i][len(s)] = memoArray[i-1][len(s)-1]
			} else {
				memoArray[i][len(s)] = 1 + min(
					memoArray[i-1][len(s)-1],
					memoArray[i-1][len(s)],
					memoArray[i][len(s)-1])
			}
		}
	}

	if len(t) >= len(s) {
		for j := 1; j < len(s); j++ {
			if t[len(t)-1] == s[j-1] {
				memoArray[len(t)][j] = memoArray[len(t)-1][j-1]
			} else {
				memoArray[len(t)][j] = 1 + min(
					memoArray[len(t)-1][j-1],
					memoArray[len(t)-1][j],
					memoArray[len(t)][j-1])
			}
		}
	}

	if s[len(s)-1] == t[len(t)-1] {
		memoArray[len(t)][len(s)] = memoArray[len(t)-1][len(s)-1]
	} else {
		memoArray[len(t)][len(s)] = 1 + min(
			memoArray[len(t)-1][len(s)-1],
			memoArray[len(t)-1][len(s)],
			memoArray[len(t)][len(s)-1])
	}

	return memoArray[len(t)][len(s)]
}
