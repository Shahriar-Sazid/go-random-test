package semioptimizedtrieed

import (
	"fmt"
)

var memoArray [100][100]float32
var replaceCost2D [26][26]float32
var insertionCost2D [26][26]float32
var deletionCostArray [26]float32

func rIndex(r rune) rune {
	i := r - 'a'
	return i
}

func si(elements []rune, i int) rune {
	if i >= 0 && i < len(elements) {
		return elements[i]
	}
	return ' '
}
func safe_cost(cost [26][26]float32, i, j rune) (result float32) {
	ii := rIndex(i)
	ji := rIndex(j)
	if ii >= 0 && ii < 26 && ji >= 0 && ji < 26 {
		return cost[ii][ji]
	}
	return 1
}

func safe_del_cost(cost [26]float32, i rune) float32 {
	ii := rIndex(i)
	if ii >= 0 && ii < 26 {
		return cost[ii]
	}
	return 1
}

func getHomophoneCost2D() [26][26]float32 {
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
		{"au", 0.6},
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
				c1i := rIndex(c1)
				c2i := rIndex(c2)
				homophoneCost[c1i][c2i] = group.cost
			}
		}
	}
	return homophoneCost
}

func getAdjacentKeyCost2D() [26][26]float32 {
	var adjacentKeyCost [26][26]float32
	for i := 0; i < len(adjacentKeyCost); i++ {
		for j := 0; j < len(adjacentKeyCost[i]); j++ {
			adjacentKeyCost[i][j] = 1
		}
	}
	isAlpha := func(c rune) bool {
		if c < 'a' || c > 'z' {
			return false
		}
		return true
	}
	qwertyKeyboard := [][]rune{
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\\'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\''},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
	}

	for i := 0; i < len(qwertyKeyboard); i++ {
		for j := 0; j < len(qwertyKeyboard[i]); j++ {
			letter := qwertyKeyboard[i][j]
			var prev rune
			if j >= 1 {
				prev = qwertyKeyboard[i][j-1]
			}

			if prev != 0 && isAlpha(prev) && isAlpha(letter) {
				p := rIndex(prev)
				l := rIndex(letter)
				adjacentKeyCost[p][l] = 0.7
				adjacentKeyCost[l][p] = 0.7
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

	homophoneCost := getHomophoneCost2D()
	adjacentKeyCost := getAdjacentKeyCost2D()

	for i := 0; i < len(replaceCost2D); i++ {
		for j := 0; j < len(replaceCost2D[i]); j++ {
			replaceCost2D[i][j] = adjacentKeyCost[i][j]
			replaceCost2D[i][j] = min(adjacentKeyCost[i][j], homophoneCost[i][j])
		}
	}

	for i := 0; i < len(insertionCost2D); i++ {
		for j := 0; j < len(insertionCost2D[i]); j++ {
			insertionCost2D[i][j] = adjacentKeyCost[i][j]
		}
	}
	insertionCost2D['s'-'a']['h'-'a'] = 0.3
	insertionCost2D['t'-'a']['h'-'a'] = 0.3
	insertionCost2D['g'-'a']['h'-'a'] = 0.3
	insertionCost2D['b'-'a']['h'-'a'] = 0.3

	for i := 0; i < len(deletionCostArray); i++ {
		deletionCostArray[i] = 1
	}
	deletionCostArray['h'-'a'] = 0.5
	fmt.Println(deletionCostArray)
}

func IncrementalED2D(s, t string, progressSoFar, steps int) float32 {
	distance := memoArray[len(t)][len(s)]
	runeS, runeT := []rune(s), []rune(t)
	for i := progressSoFar; i < progressSoFar+steps; i++ {
		distance = incrementalED2D(runeS[:min(len(runeS), i+1)], runeT[:min(len(runeT), i+1)])
	}
	return distance
}

func incrementalED2D(s, t []rune) float32 {
	if len(s) >= len(t) {
		for i := 1; i < len(t); i++ {
			if s[len(s)-1] == t[i-1] {
				memoArray[i][len(s)] = memoArray[i-1][len(s)-1]
			} else {
				rc := memoArray[i-1][len(s)-1] + safe_cost(replaceCost2D, s[len(s)-1], t[i-1])
				ic := memoArray[i-1][len(s)] + safe_cost(insertionCost2D, si(t, i-2), si(t, i-1))
				dc := memoArray[i][len(s)-1] + safe_del_cost(deletionCostArray, s[len(s)-1])
				memoArray[i][len(s)] = min(rc, ic, dc)
			}
		}
	}

	if len(t) >= len(s) {
		for j := 1; j < len(s); j++ {
			if t[len(t)-1] == s[j-1] {
				memoArray[len(t)][j] = memoArray[len(t)-1][j-1]
			} else {
				rc := memoArray[len(t)-1][j-1] + safe_cost(replaceCost2D, s[j-1], t[len(t)-1])
				ic := memoArray[len(t)-1][j] + safe_cost(insertionCost2D, t[len(t)-2], t[len(t)-1])
				dc := memoArray[len(t)][j-1] + safe_del_cost(deletionCostArray, s[j-1])
				memoArray[len(t)][j] = min(rc, ic, dc)
			}
		}
	}

	if s[len(s)-1] == t[len(t)-1] {
		memoArray[len(t)][len(s)] = memoArray[len(t)-1][len(s)-1]
	} else {
		rc := memoArray[len(t)-1][len(s)-1] + safe_cost(replaceCost2D, s[len(s)-1], t[len(t)-1])
		ic := memoArray[len(t)-1][len(s)] + safe_cost(insertionCost2D, si(t, len(t)-2), si(t, len(t)-1))
		dc := memoArray[len(t)][len(s)-1] + safe_del_cost(deletionCostArray, s[len(s)-1])
		memoArray[len(t)][len(s)] = min(rc, ic, dc)
	}

	return memoArray[len(t)][len(s)]
}
func IncrementalMatchRatio2D(s, t string, progressSoFar, steps int) float32 {
	distance := IncrementalED2D(s, t, progressSoFar, steps)
	maxLen := progressSoFar + steps

	ratio := 1 - (distance / float32(maxLen))
	// fmt.Printf("%d. match ratio of %s,%s is %f\n", counter, s[:min(len(s), progressSoFar+steps)], t[:min(len(t), progressSoFar+steps)], ratio)
	return ratio
}
