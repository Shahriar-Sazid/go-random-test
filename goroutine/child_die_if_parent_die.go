package routine

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generateRandomLine() string {
	// Define a pool of characters to choose from for the random line.
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 20 // You can adjust the length of the random line as needed.

	// Generate a random line by selecting characters from the pool.
	rand.Seed(time.Now().UnixNano())
	randomLine := make([]byte, length)
	for i := range randomLine {
		randomLine[i] = characters[rand.Intn(len(characters))]
	}

	return string(randomLine)
}

func writeLineToFile() (string, error) {
	time.Sleep(time.Second)
	// Generate a random filename based on the current timestamp and a random number.
	rand.Seed(time.Now().UnixNano())
	filename := fmt.Sprintf("output_%d.txt", rand.Intn(1000))

	// Generate a random line.
	line := generateRandomLine()

	// Open a file for writing. If the file doesn't exist, create it; if it exists, truncate it.
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close() // Ensure the file is closed when we're done with it.

	// Write the random line to the file.
	_, err = file.WriteString(line)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func TestChildDieIfParentDie() {
	for i := 0; i < 10; i++ {
		go writeLineToFile()
	}
}
