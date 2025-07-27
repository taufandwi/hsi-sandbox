## Key Concepts - Packages

* Go programs are organized into **packages**.
* A package is a collection of Go source files in the same directory that are compiled together.
* **`package main`**: Special package name indicating an executable program. The `main` function in this package is the entry point.
* **Importing Packages:** Use the `import` keyword.
    ```go
    import (
        "fmt"       // Standard library package
        "math/rand" // Sub-package from math
        "net/http"  // For HTTP clients and servers
    )
    ```
* **Exported Names:** If a name (variable, function, type, etc.) starts with a capital letter, it is exported from the package and can be accessed by other packages. Names starting with a lowercase letter are private to the package.
    * `fmt.Println` (Println is exported)
    * `math.Pi` (Pi is exported)
* **Creating Your Own Packages:**
    1.  Create a new directory (e.g., `myutils`).
    2.  Inside `myutils`, create Go files (e.g., `stringutils.go`) with `package myutils` at the top.
    3.  Define exported functions/variables in these files.
    4.  In your `main` package (or another package), import it using its path (e.g., `import "yourproject/myutils"` if using Go Modules).

---

## Key Concepts - Go Modules
* **Go Modules**: A dependency management system introduced in Go 1.11.
* It allows you to manage dependencies and versions of your Go projects.
* **Creating a Module:**
    1. Navigate to your project directory.
    2. Run `go mod init <module-name>` to create a new module.
    3. Add dependencies using `go get <package>`.
  4. Use `go mod tidy` to clean up unused dependencies.
  5. Use `go build` to compile your module, which will automatically download dependencies.
  6. Use `go run .` to run your module directly.
  7. Use `go test` to run tests in your module.
  8. Use `go list` to list all packages in your module.
  9. Use `go mod verify` to check that dependencies have not been modified.
  10. Use `go mod edit` to edit the `go.mod` file directly.
  11. Use `go mod graph` to visualize the dependency graph of your module.
  12. Use `go mod why` to understand why a dependency is included in your module.

--

## Latihan

1. Dari Project minggu ke-3, buatlah sebuah package baru yang berisi fungsi-fungsi yang berada pada file `main.go` yang telah Anda buat. 
Pastikan fungsi-fungsi tersebut dapat diakses dari package lain.

Setelah menyelesaikan level ini, Anda akan memahami cara mendefinisikan dan menggunakan interface di Go.

Selamat belajar!