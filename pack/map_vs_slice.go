package pack

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"time"
)

func TestMapVSSlice() {
	// Create a slice of 1000 unique random numbers
	slice := generateRandomSlice(40)

	// Create a map from the slice where keys are slice elements and values are indices
	indexMap := createIndexMap(slice)

	var start, end time.Time
	var elapsedTime int64

	start = time.Now()
	for _, element := range slice {
		// Find element index using loop
		_ = findIndexUsingLoop(slice, element)

	}
	end = time.Now()

	elapsedTime = end.Sub(start).Nanoseconds()
	fmt.Printf("loop took %d ns to execute\n", elapsedTime)

	start = time.Now()
	for _, element := range slice {
		// Find element index using map
		_ = findIndexUsingMap(indexMap, element)
	}
	end = time.Now()

	elapsedTime = end.Sub(start).Nanoseconds()
	fmt.Printf("map took %d ns to execute\n", elapsedTime)

	start = time.Now()
	for _, element := range slice {
		// Find element index using map
		_ = findIndexUsingBinarySearch(slice, element)
	}
	end = time.Now()

	elapsedTime = end.Sub(start).Nanoseconds()
	fmt.Printf("binary search took %d ns to execute\n", elapsedTime)
}

// generateRandomSlice generates a slice of specified length with unique random numbers
func generateRandomSlice(length int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)
	used := make(map[int]bool)

	for i := 0; i < length; {
		num := rand.Intn(length * 10) // Assuming range of random numbers is 0 to 10 times length
		if !used[num] {
			slice[i] = num
			used[num] = true
			i++
		}
	}
	sort.Ints(slice)
	return slice
}

// createIndexMap creates a map from a slice where keys are elements and values are indices
func createIndexMap(slice []int) map[int]int {
	indexMap := make(map[int]int)
	for i, num := range slice {
		indexMap[num] = i
	}
	return indexMap
}

// findIndexUsingLoop finds the index of an element in a slice using a loop
func findIndexUsingLoop(slice []int, element int) int {
	for i, num := range slice {
		if num == element {
			return i
		}
	}
	return -1 // element not found
}

// findIndexUsingLoop finds the index of an element in a slice using a loop
func findIndexUsingBinarySearch(slice []int, element int) int {
	idx, ok := slices.BinarySearch(slice, element)
	if ok {
		return idx
	}
	return -1 // element not found
}

// findIndexUsingMap finds the index of an element in a slice using a pre-created map
func findIndexUsingMap(indexMap map[int]int, element int) int {
	index, found := indexMap[element]
	if found {
		return index
	}
	return -1 // element not found
}
