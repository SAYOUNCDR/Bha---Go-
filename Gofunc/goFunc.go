package main

import (
	"errors"
	"fmt"
)

// Division function that returns result and error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Case 1: Use both result and error
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Case 2: Only interested in error, not result
	_, err2 := divide(5, 0)
	if err2 != nil {
		fmt.Println("Error (ignored result):", err2)
	}

	// Case 3: Only interested in result (not best practice)
	resOnly, _ := divide(8, 2) // Ignore error (not recommended in real use)
	fmt.Println("Result (ignored error):", resOnly)
}
 