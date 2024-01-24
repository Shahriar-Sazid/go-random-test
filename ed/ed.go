package ed

import (
	"encoding/csv"
	"fmt"
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

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
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

	// Print CSV contents

	memo := make(map[key]int)
	start := time.Now()
	for _, record := range records {
		edRecursionDP(record[0], record[1], memo)
	}
	end := time.Now()

	elapsedTime := end.Sub(start).Milliseconds()
	fmt.Printf("function took %d ms to execute\n", elapsedTime)

	start = time.Now()
	for _, record := range records {
		editDistanceDP(record[0], record[1])
	}
	end = time.Now()

	elapsedTime = end.Sub(start).Milliseconds()
	fmt.Printf("function took %d ms to execute\n", elapsedTime)

	lev := metrics.NewLevenshtein()
	lev.CaseSensitive = true
	lev.InsertCost = 1
	lev.ReplaceCost = 1
	lev.DeleteCost = 1

	start = time.Now()
	for _, record := range records {
		strutil.Similarity(record[0], record[1], lev)
	}
	end = time.Now()
	elapsedTime = end.Sub(start).Milliseconds()
	fmt.Printf("function took %d ms to execute\n", elapsedTime)
}
