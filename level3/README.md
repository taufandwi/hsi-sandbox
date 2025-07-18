## Key Concepts - Structs

* A **struct** is a composite data type that groups together zero or more named values (fields) of arbitrary types.
* They are useful for creating custom data structures.
* Similar to classes in object-oriented languages, but Go is not strictly OO (no inheritance in the traditional sense, but composition via embedding).

* **Defining a Struct:**
    ```go
    type Person struct {
        Name string
        Age  int
    }
    ```

* **Creating Instances:**
    ```go
    p1 := Person{"Alice", 30}
    p2 := Person{Name: "Bob", Age: 25}
    p3 := Person{Name: "Charlie"} // Age will be 0 (zero value for int)
    var p4 Person // All fields zero-valued (Name="", Age=0)
    ```

* **Accessing Fields:** Use the dot `.` operator.
    ```go
    fmt.Println(p1.Name) // Alice
    p2.Age = 26
    ```

* **Pointers to Structs:**
    ```go
    ptrP := &p1
    ptrP.Age = 31 // Go automatically dereferences: (*ptrP).Age = 31
    fmt.Println(p1.Age) // 31
    ```

* **Methods on Structs:** Functions with a special receiver argument.
    ```go
    // Method with a value receiver
    func (p Person) Greet() {
        fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
    }

    // Method with a pointer receiver (to modify the struct)
    func (p *Person) Birthday() {
        p.Age++
    }

    // Usage:
    // p1.Greet()
    // p1.Birthday()
    ```

---

## Key Concepts - Interfaces
* An **interface** is a type that specifies a contract of methods that a type must implement.
* Interfaces allow for polymorphism, enabling different types to be treated uniformly based on shared behavior.
* An interface is defined using the `type` keyword followed by the interface name and its method signatures.
* **Defining an Interface:**
    ```go
    package main

    import "fmt"
    
    // Define an interface
    type Shape interface {
        Area() float64       // Method signature: Area returns a float64
        Perimeter() float64  // Method signature: Perimeter returns a float64
    }
    
    // Define a struct type Circle
    type Circle struct {
        Radius float64
    }
    
    // Implement the Area method for Circle
    func (c Circle) Area() float64 {
        return 3.14159 * c.Radius * c.Radius
    }
    
    // Implement the Perimeter method for Circle
    func (c Circle) Perimeter() float64 {
        return 2 * 3.14159 * c.Radius
    }
    
    // Define a struct type Rectangle
    type Rectangle struct {
        Width  float64
        Height float64
    }
    
    // Implement the Area method for Rectangle
    func (r Rectangle) Area() float64 {
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

        printShapeInfo(c) // Circle satisfies Shape, so it can be passed
        printShapeInfo(r) // Rectangle satisfies Shape, so it can be passed

        // You can also declare a variable of interface type
        var s Shape
        s = c // 's' now holds a Circle value
        fmt.Println("Area of s (Circle):", s.Area())

        s = r // 's' now holds a Rectangle value
        fmt.Println("Area of s (Rectangle):", s.Area())
    }
    ```




## Latihan

1. **Model Buku Perpustakaan**: Buat struct untuk buku (`judul buku`, `jumlah ketersediaan`), buatlah method untuk apabila buku dipinjam dan mengembalikannya, serta menampilkan informasi seluruh judul buku dan ketersediaanya.


Selamat belajar!