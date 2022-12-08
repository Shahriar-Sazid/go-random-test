package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func FileTest() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 0
	inputFile, err := os.Create("input.csv") //create the input.csv file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		inputFile.Close()
	}()

	csvwriter := csv.NewWriter(inputFile)
	defer func() {
		csvwriter.Flush()
	}()
	fmt.Println("How many records ?")
	fmt.Scanln(&n)
	fmt.Println("Enter the records")
	var lines [][]string
	for i := 0; i < n; i++ {
		scanner.Scan()
		text := scanner.Text()
		lines = append(lines, []string{text})

	}
	err = csvwriter.WriteAll(lines)
	if err != nil {
		return
	}
}
