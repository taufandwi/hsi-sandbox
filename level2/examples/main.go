package main

import (
	"fmt"
	"runtime/debug"
)

// divide performs division and might panic if denominator is zero.
func divide(a, b int) int {
	// This deferred function will be executed if a panic occurs in `divide`.
	// It attempts to recover the panic.
	defer func() {
		if r := recover(); r != nil {
			// A panic was recovered!
			fmt.Printf("\n--- Recovered from panic in divide() ---\n")
			fmt.Printf("Recovered value: %v\n", r)
			// Print stack trace for debugging purposes
			fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
			fmt.Printf("--- End of recover in divide() ---\n\n")
		}
	}()
	// sama seperti try catch di java atau bahasa lain

	fmt.Printf("Attempting to divide %d by %d...\n", a, b)
	if b == 0 {
		// Panic if the denominator is zero.
		// This is an unrecoverable error for this specific division operation.
		panic("denominator cannot be zero") // fmt.errorf("denominator cannot be zero")

	}
	result := a / b
	fmt.Printf("Division successful: %d / %d = %d\n", a, b, result)
	return result
}

func main() {
	fmt.Println("Starting main function.")

	// --- Scenario 1: Panic and Recover within the same call stack ---
	fmt.Println("\n--- Scenario 1: Panic and Recover in divide() ---")
	result1 := divide(10, 2)
	fmt.Printf("Result 1: %d\n", result1) // This line executes normally

	result2 := divide(10, 0) // This call will panic, but `recover` in `divide` will handle it.
	fmt.Printf("Result 2: %d\n", result2)
}
