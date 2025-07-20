package main

import (
	"fmt"
)

// Shape is an interface that defines methods for shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// Create a slice of Shape
	shapes := []Shape{rect, circ}

	// Iterate over shapes and print area and perimeter
	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f, Perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}
}

// Exercise:
// 1. Create a new struct called Triangle that implements the Shape interface.
// 2. Implement the Area and Perimeter methods for Triangle.
// 3. Add an instance of Triangle to the shapes slice and print its area and perimeter.
