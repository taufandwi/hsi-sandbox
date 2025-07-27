package main

import "fmt"

func main() {
	inputText := "aabbaa" // Change this to test different strings

	n := len(inputText)

	isPalindrome := true

	for i := 0; i < n/2; i++ {
		if inputText[i] != inputText[n-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println(inputText, "is a palindrome text")
	} else {
		fmt.Println(inputText, "is not a palindrome text")
	}
}
