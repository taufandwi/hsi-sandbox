package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64      // Method signature: Area returns a float64
	Perimeter() float64 // Method signature: Perimeter returns a float64
	//GetNumber(sisi int) int
}

// Define a struct type Circle
type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * (c.Radius * c.Radius)
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	// check run time
	fmt.Println("this is a Circle method area")
	return 3.14159 * c.Radius * c.Radius
}

// Define a struct type Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	fmt.Println("this is a Rectangle method area")
	return r.Width * r.Height
}

// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// A function that accepts any type that satisfies the Shape interface
func printShapeInfo(s Shape) {
	fmt.Printf("Shape Info:\n  Area: %.2f\n  Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Printf("Circle Area: %.2f\n", c.Area())
	fmt.Printf("Rectangle Area: %.2f\n", r.Area())

	//printShapeInfo(c) // Circle satisfies Shape, so it can be passed
	//printShapeInfo(r) // Rectangle satisfies Shape, so it can be passed

	// You can also declare a variable of interface type
	var s Shape
	s = c // 's' now holds a Circle value
	fmt.Println("Area of s (Circle):", s.Area())

	s = r // 's' now holds a Rectangle value
	fmt.Println("Area of s (Rectangle):", s.Area())
}
