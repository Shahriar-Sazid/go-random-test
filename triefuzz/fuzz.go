package triefuzz

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Shahriar-Sazid/go-random-test/trie"
)

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

	words := make([]string, 30000)
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
	t := trie.NewTrie()

	for _, word := range words {
		t.Insert(word)
	}

	word := "siddhirhgonj"
	startTime := time.Now()
	_ = t.FuzzySearch(word)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Microseconds()
	fmt.Printf("trie fuzz took %d us to search %s\n", elapsedTime, word)
	// for _, result := range results {
	// 	fmt.Println(result.Word, result.Token, result.Ratio)
	// }
}
