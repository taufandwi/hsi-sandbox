package main

import (
	"fmt"
	"github.com/taufandwi/hsi-sandbox/tree/week_one/level4/etc"
)

func main() {

	fmt.Println("Hello World")

	fmt.Println("This is a test for Go module")
	fmt.Println("This is a test for Go module with custom time package")
	fmt.Println("Current time is:", etc.GetCurrentTime())

	fmt.Println("This is a test for Go module with UUID generator")
	fmt.Println("Generated UUID:", etc.GenerateUUID())

}
