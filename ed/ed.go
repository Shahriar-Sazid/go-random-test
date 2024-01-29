package ed

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

type key struct {
	s1 string
	s2 string
}

func edRecursionDP(str1, str2 string, memo map[key]int) int {
	if len(str1) == 0 {
		return len(str2)
	} else if len(str2) == 0 {
		return len(str1)
	}

	k := key{str1, str2}
	if val, ok := memo[k]; ok {
		return val
	}

	var result int
	if str1[len(str1)-1] == str2[len(str2)-1] {
		result = edRecursionDP(str1[:len(str1)-1], str2[:len(str2)-1], memo)
	} else {
		deletion := edRecursionDP(str1[:len(str1)-1], str2, memo)
		insertion := edRecursionDP(str1, str2[:len(str2)-1], memo)
		substitution := edRecursionDP(str1[:len(str1)-1], str2[:len(str2)-1], memo)
		result = min(deletion, insertion, substitution) + 1
	}

	memo[k] = result
	return result
}

func editDistanceRecursive(str1, str2 string, m, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if str1[m-1] == str2[n-1] {
		return editDistanceRecursive(str1, str2, m-1, n-1)
	}

	insertCost := editDistanceRecursive(str1, str2, m, n-1)
	deleteCost := editDistanceRecursive(str1, str2, m-1, n)
	replaceCost := editDistanceRecursive(str1, str2, m-1, n-1)

	return 1 + min(insertCost, deleteCost, replaceCost)
}

func editDistanceRecursiveWithMemo(str1, str2 string, m, n int, memo map[[2]int]int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if val, found := memo[[2]int{m, n}]; found {
		return val
	}

	var result int
	if str1[m-1] == str2[n-1] {
		result = editDistanceRecursiveWithMemo(str1, str2, m-1, n-1, memo)
	} else {
		insertCost := editDistanceRecursiveWithMemo(str1, str2, m, n-1, memo)
		deleteCost := editDistanceRecursiveWithMemo(str1, str2, m-1, n, memo)
		replaceCost := editDistanceRecursiveWithMemo(str1, str2, m-1, n-1, memo)

		result = 1 + min(insertCost, deleteCost, replaceCost)
	}

	memo[[2]int{m, n}] = result
	return result
}

func editDistanceRecursive1(str1, str2 string) int {
	memo := make(map[[2]int]int)
	return editDistanceRecursiveWithMemo(str1, str2, len(str1), len(str2), memo)
}

func editDistanceDP(str1, str2 string) int {
	m := len(str1)
	n := len(str2)

	// Create a 2D slice to store the edit distances
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the base cases
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	// The bottom-right cell contains the edit distance
	return dp[m][n]
}

type st struct {
	s string
	t string
}
type cost struct {
	memo map[st]float32
}

func (c cost) get(s string, t string) float32 {
	cost, ok := c.memo[st{s, t}]
	if ok {
		return cost
	}

	if len(t) == 0 {
		return float32(len(s))
	}

	if len(s) == 0 {
		return float32(len(t))
	}

	return cost
}

func (c cost) set(s string, t string, cost float32) {
	c.memo[st{s, t}] = cost
}

var cst cost

func ed(s string, t string) float32 {
	sl := len(s)
	tl := len(t)
	d := float32(math.Abs(float64(sl - tl)))
	m := min(sl, tl)

	if sl != tl {
		return cst.get(s[:m], t[:m]) + d
	}

	for i := 0; i < m-1; i++ {
		var distance float32
		distance = min(
			cst.get(s, t[:i]),
			cst.get(s[:len(s)-1], t[:i]),
			cst.get(s[:len(s)-1], t[:i+1]),
		)
		if s[len(s)-1] != t[i] {
			distance++
		}
		cst.set(s, t[:i+1], distance)

		distance = min(
			cst.get(s[:i], t),
			cst.get(s[:i], t[:len(t)-1]),
			cst.get(s[:i+1], t[:len(t)-1]),
		)
		if s[i] != t[len(t)-1] {
			distance++
		}
		cst.set(s[:i+1], t, distance)
	}
	distance := min(
		cst.get(s[:sl-1], t[:tl-1]),
		cst.get(s, t[:tl-1]),
		cst.get(s[:sl-1], t),
	)
	if s[sl-1] != t[tl-1] {
		distance++
	}
	cst.set(s, t, distance)

	return distance
}

func TestED() {
	file, err := os.Open("ed_sample.csv")
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

	for i := 0; i < len(records); i++ {
		records[i][0] = records[i][0] + records[i][0] + records[i][0]
		records[i][1] = records[i][1] + records[i][1] + records[i][1]
	}

	// Print CSV contents

	memo := make(map[key]int)
	start := time.Now()
	for _, record := range records {
		edRecursionDP(record[0], record[1], memo)
	}
	end := time.Now()

	elapsedTime := end.Sub(start).Milliseconds()
	fmt.Printf("recursion took %d ms to execute\n", elapsedTime)

	start = time.Now()
	for _, record := range records {
		editDistanceDP(record[0], record[1])
	}
	end = time.Now()

	elapsedTime = end.Sub(start).Milliseconds()
	fmt.Printf("loop took %d ms to execute\n", elapsedTime)

	lev := metrics.NewLevenshtein()
	lev.CaseSensitive = true
	lev.InsertCost = 1
	lev.ReplaceCost = 1
	lev.DeleteCost = 1

	cst = cost{make(map[st]float32, 10000)}
	start = time.Now()
	for _, record := range records {
		strutil.Similarity(record[0], record[1], lev)
	}
	end = time.Now()
	elapsedTime = end.Sub(start).Milliseconds()
	fmt.Printf("lib took %d ms to execute\n", elapsedTime)

	start = time.Now()
	for _, record := range records {
		minLen := min(len(record[0]), len(record[1]))
		for i := 0; i < minLen; i++ {
			ed(record[0][:i+1], record[1][:i+1])
		}
	}
	end = time.Now()
	elapsedTime = end.Sub(start).Milliseconds()
	fmt.Printf("diagonal function took %d ms to execute\n", elapsedTime)

	start = time.Now()
	for _, record := range records {
		minLen := min(len(record[0]), len(record[1]))
		for i := 0; i < minLen; i++ {
			strutil.Similarity(record[0][:i+1], record[1][:i+1], lev)
		}
	}
	end = time.Now()
	elapsedTime = end.Sub(start).Milliseconds()
	fmt.Printf("lib1 function took %d ms to execute\n", elapsedTime)

	source := "bagerhat"
	target := "bosurhat"

	cst = cost{make(map[st]float32)}
	var distance float32
	for i := 0; i < len(source); i++ {
		distance = ed(source[:i+1], target[:i+1])
	}
	fmt.Println(distance)
}
