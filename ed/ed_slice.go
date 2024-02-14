package ed

type FuzzyResult struct {
	Word  string
	Token string
	Ratio float32
}

var memoArray [100][100]float32

var replaceCost []float32
var insertionCost []float32
var deletionCost []float32

func runeIndex(r rune) int {
	i := r - 'a'
	return int(i)
}

func twoToOneD(i, j rune, rowSize int) int {
	return runeIndex(i)*rowSize + runeIndex(j)
}

func safe_index[T any](s []T, index int, standard T) T {
	if index < 0 || index >= len(s) {
		return standard
	}
	return s[index]
}

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
				c1i := runeIndex(c1)
				c2i := runeIndex(c2)
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
				p := runeIndex(prev)
				l := runeIndex(letter)
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
	replaceCost = make([]float32, 26*26)
	insertionCost = make([]float32, 26*26)
	deletionCost = make([]float32, 26)

	homophoneCost := getHomophoneCost()
	adjacentKeyCost := getAdjacentKeyCost()

	for i := 0; i < len(adjacentKeyCost); i++ {
		for j := 0; j < len(adjacentKeyCost[i]); j++ {
			replaceCost[i*26+j] = adjacentKeyCost[i][j]
			replaceCost[i*26+j] = min(adjacentKeyCost[i][j], homophoneCost[i][j])
		}
	}

	for i := 0; i < len(adjacentKeyCost); i++ {
		for j := 0; j < len(adjacentKeyCost[i]); j++ {
			insertionCost[i*26+j] = adjacentKeyCost[i][j]
		}
	}
	insertionCost[('s'-'a')*26+('h'-'a')] = 0.3
	insertionCost[('t'-'a')*26+('h'-'a')] = 0.3
	insertionCost[('g'-'a')*26+('h'-'a')] = 0.3
	insertionCost[('b'-'a')*26+('h'-'a')] = 0.3

	for i := 0; i < len(deletionCost); i++ {
		deletionCost[i] = 1
	}
	deletionCost['h'-'a'] = 0.5
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
				rc := memoArray[i-1][len(s)-1] + safe_index(replaceCost, twoToOneD(s[len(s)-1], t[i-1], 26), 1)
				ic := memoArray[i-1][len(s)] + safe_index(insertionCost, twoToOneD(safe_index(t, i-2, ' '), safe_index(t, i-1, ' '), 26), 1)
				dc := memoArray[i][len(s)-1] + safe_index(deletionCost, runeIndex(s[len(s)-1]), 1)
				memoArray[i][len(s)] = min(rc, ic, dc)
			}
		}
	}

	if len(t) >= len(s) {
		for j := 1; j < len(s); j++ {
			if t[len(t)-1] == s[j-1] {
				memoArray[len(t)][j] = memoArray[len(t)-1][j-1]
			} else {
				rc := memoArray[len(t)-1][j-1] + safe_index(replaceCost, twoToOneD(s[j-1], t[len(t)-1], 26), 1)
				ic := memoArray[len(t)-1][j] + safe_index(insertionCost, twoToOneD(t[len(t)-2], t[len(t)-1], 26), 1)
				dc := memoArray[len(t)][j-1] + safe_index(deletionCost, runeIndex(s[j-1]), 1)
				memoArray[len(t)][j] = min(rc, ic, dc)
			}
		}
	}

	if s[len(s)-1] == t[len(t)-1] {
		memoArray[len(t)][len(s)] = memoArray[len(t)-1][len(s)-1]
	} else {
		rc := memoArray[len(t)-1][len(s)-1] + safe_index(replaceCost, twoToOneD(s[len(s)-1], t[len(t)-1], 26), 1)
		ic := memoArray[len(t)-1][len(s)] + safe_index(insertionCost, twoToOneD(safe_index(t, len(t)-2, ' '), safe_index(t, len(t)-1, ' '), 26), 1)
		dc := memoArray[len(t)][len(s)-1] + safe_index(deletionCost, runeIndex(s[len(s)-1]), 1)
		memoArray[len(t)][len(s)] = min(rc, ic, dc)
	}

	return memoArray[len(t)][len(s)]
}
