package pack

import (
	"fmt"
	"math/big"
)

func GCDTest() {
	// Replace these values with the integers for which you want to calculate the GCD
	numbers := []int64{48, 18, 36, 72}

	// Create a big.Int to hold the GCD
	gcd := big.NewInt(numbers[0])

	// Calculate the GCD of all numbers
	for i := 1; i < len(numbers); i++ {
		bigNum := big.NewInt(numbers[i])
		gcd.GCD(nil, nil, gcd, bigNum)
	}

	fmt.Printf("GCD of %v is %s\n", numbers, gcd)
}
