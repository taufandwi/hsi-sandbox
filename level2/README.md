## Basic Syntax - Functions

* Functions are fundamental building blocks.
* **Definition:**
    ```go
    func functionName(param1 type1, param2 type2) returnType {
        // function body
        return value
    }
    ```
* **Multiple Return Values:** Go functions can return multiple values. This is often used to return a result and an error.
    ```go
    func divide(a, b int) (int, error) {
        if b == 0 {
            return 0, fmt.Errorf("cannot divide by zero")
        }
        return a / b, nil // nil means no error
    }
    ```

* **Example:**
    ```go
    package main

    import "fmt"

    // Function to add two integers
    func add(x int, y int) int {
        return x + y
    }

    // Function returning multiple values
    func swap(a, b string) (string, string) {
        return b, a
    }

    func main() {
        result := add(5, 3)
        fmt.Println("5 + 3 =", result)

        x, y := "hello", "world"
        fmt.Println("Before swap:", x, y)
        x, y = swap(x, y)
        fmt.Println("After swap:", x, y)

        quotient, err := divide(10, 2)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("10 / 2 =", quotient)
        }

        _, err = divide(10, 0) // Using blank identifier _ to ignore quotient
        if err != nil {
            fmt.Println("Error:", err)
        }
    }
    ```

---

## Key Concepts - Error Handling

* Go has a built-in `error` type, which is an interface.
* By convention, functions that can fail return an `error` as their last return value.
* A `nil` error value indicates success; a non-`nil` value indicates failure.
* **Idiomatic Error Handling:** Always check the error value returned by a function.

* **Example:**
    ```go
    package main

    import (
        "errors" // For creating simple error messages
        "fmt"
        "strconv" // String conversion, can produce errors
    )

    // A function that might return an error
    func mightFail(shouldFail bool) (string, error) {
        if shouldFail {
            return "", errors.New("something went wrong")
        }
        return "success!", nil
    }

    func main() {
        result, err := mightFail(false)
        if err != nil {
            fmt.Println("Error occurred:", err)
        } else {
            fmt.Println("Result:", result)
        }

        result, err = mightFail(true)
        if err != nil {
            fmt.Println("Error occurred:", err)
        } else {
            fmt.Println("Result:", result)
        }

        // Example from standard library
        numStr := "123"
        num, err := strconv.Atoi(numStr) // Atoi converts string to int
        if err != nil {
            fmt.Printf("Could not convert '%s': %v\n", numStr, err)
        } else {
            fmt.Printf("Converted '%s' to %d\n", numStr, num)
        }

        numStr = "abc"
        num, err = strconv.Atoi(numStr)
        if err != nil {
            fmt.Printf("Could not convert '%s': %v\n", numStr, err)
        } else {
            fmt.Printf("Converted '%s' to %d\n", numStr, num)
        }
    }
    ```
* **Panic and Recover:** For truly exceptional errors, Go has `panic` (stops normal execution) and `recover` (regains control of a panicking goroutine). Use sparingly.
  * **Example:**
      ```go
      package main
  
      import "fmt"
  
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
    
          fmt.Printf("Attempting to divide %d by %d...\n", a, b)
          if b == 0 {
              // Panic if the denominator is zero.
              // This is an unrecoverable error for this specific division operation.
              panic("denominator cannot be zero")
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
    ```
---

## Slide 10: Key Concepts - Pointers

* A **pointer** holds the memory address of a variable.
* `*T` is the type of a pointer to a `T` value. (e.g., `*int` is a pointer to an `int`)
* The `&` operator generates a pointer to its operand.
    ```go
    i := 42
    p := &i // p now holds the memory address of i; p is of type *int
    ```
* The `*` operator (dereferencing) accesses the underlying value through the pointer.
    ```go
    fmt.Println(*p) // reads i through the pointer p (outputs 42)
    *p = 21         // sets i through the pointer p (i is now 21)
    ```
* **Why use pointers?**
    * To allow functions to modify the value of a variable passed as an argument.
    * For efficiency with large data structures (avoids copying the entire structure).
    * The zero value for a pointer is `nil`.

* **Example:**
    ```go
    package main

    import "fmt"

    func zeroval(ival int) {
        ival = 0 // This changes a copy of ival
    }

    func zeroptr(iptr *int) {
        *iptr = 0 // This changes the value at the memory address iptr points to
    }

    func main() {
        i := 1
        fmt.Println("initial:", i) // 1

        zeroval(i)
        fmt.Println("zeroval:", i) // 1 (i is not changed)

        zeroptr(&i) // Pass the memory address of i
        fmt.Println("zeroptr:", i) // 0 (i is changed)

        fmt.Println("pointer address:", &i) // Prints the memory address
    }
    ```

---

## Latihan

buatlah fungsi-fungsi berikut:
1. **Fungsi Swap:** Buat fungsi yang menerima dua string tanpa `return` dan mengembalikannya dalam urutan terbalik.(gunakan pointer untuk mengubah nilai asli)
2. **Fungsi Pencarian Bilangan Terbesar:** Buat fungsi yang menerima slice dari bilangan bulat dan mengembalikan bilangan terbesar dalam slice tersebut.
3. **Fungsi Pencarian Bilangan Genap:** Buat fungsi yang menerima slice dari bilangan bulat dan mengembalikan slice baru yang hanya berisi bilangan genap dari slice input. kembalikan error jika tidak ada bilangan genap yang ditemukan.
4. **Fungsi Pencarian Bilangan Ganjil:** Buat fungsi yang menerima slice dari bilangan bulat dan mengembalikan slice baru yang hanya berisi bilangan ganjil dari slice input. kembalikan error jika tidak ada bilangan ganjil yang ditemukan.

Selamat belajar dan praktikkan materi ini!