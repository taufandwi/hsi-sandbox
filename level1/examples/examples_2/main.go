package main

import "fmt"

func main() {
	var greeting string = "Hello from a variable!"
	fmt.Println(greeting)

	var a, b int = 10, 20
	fmt.Println("Sum:", a+b)

	// Short declaration
	isEnabled := true
	name := "Go Developer"
	fmt.Println("Status:", isEnabled, "User:", name)

	// part 2
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// With a short statement
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
