package panicdefer

import "fmt"

func TestPanicRecover() (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 20 // Set a default value
		}
	}()

	// Simulate a panic condition
	if true {
		panic("panic occurred")
	}

	// This line won't be reached due to the panic
	return 42
}
