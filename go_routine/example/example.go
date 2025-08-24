package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello!")
		time.Sleep(100 * time.Millisecond)
	}
}

func sayWorld() {
	for i := 0; i < 3; i++ {
		fmt.Println("World!")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting the program.")

	sayHello() // Start a new goroutine for sayHello()
	sayWorld() // Start a new goroutine for sayWorld()

	//go sayHello() // Start a new goroutine for sayHello()
	//go sayWorld() // Start a new goroutine for sayWorld()

	// Wait for a little while to give the goroutines time to run
	//time.Sleep(500 * time.Millisecond)

	// What happens now?
	fmt.Println("Program finished.")
}
