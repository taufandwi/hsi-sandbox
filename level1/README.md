# Level 1: Sintaks Dasar dan Tipe Data di Go

Selamat datang di Level 1 **HSI - Golang Sandbox**! Pada level ini, Anda akan mempelajari dasar-dasar sintaks dan tipe data di Go, termasuk variabel, konstanta, dan struktur kontrol.

## Materi

1. **Variabel**: Cara mendeklarasikan dan menggunakan variabel di Go.
2. **Konstanta**: Pengertian dan cara mendefinisikan konstanta.
3. **Struktur Kontrol Dasar**: Penggunaan `if`, `for`, dan `switch`.

## Basic Syntax - Variables & Data Types
* **Variable Declaration:**
    * `var name type` (e.g., `var age int`)
    * `var name type = value` (e.g., `var message string = "Hi"`)
    * `name := value` (short declaration, Go infers the type, only usable inside functions)

* **Common Data Types:**
    * `int`, `int8`, `int16`, `int32`, `int64` (integers)
    * `uint`, `uint8`, ... (unsigned integers)
    * `float32`, `float64` (floating-point numbers)
    * `string` (text)
    * `bool` (`true` or `false`)
    * `rune` (a character, an alias for `int32`)
    * `byte` (an alias for `uint8`)
---

## Basic Syntax - Control Flow: `if/else`
* Used for conditional execution.
* **Syntax:**
    * No parentheses `()` around the condition.
    * Braces `{}` are always required, even for single-line statements.

* **Forms:**
    * `if condition { ... }`
    * `if condition { ... } else { ... }`
    * `if condition1 { ... } else if condition2 { ... } else { ... }`
    * `if` can also include a short statement executed before the condition.
---

## Basic Syntax - Control Flow: Looping with `for`

* Go has only one looping construct: the `for` loop. It can be used in several ways.

1.  **Complete `for` statement (C-style):**
    ```go
    for initialization; condition; post_statement {
        // loop body
    }
    ```
    **Example:**
    ```go
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration:", i)
    }
    ```
    
2.  **Condition-only `for` (like `while`):**
    ```go
    sum := 1
    for sum < 100 { // while sum < 100
        sum += sum
    }
    fmt.Println("Sum:", sum)
    ```

3.  **Infinite loop (with `break` or `return`):**
    ```go
    // for {
    //     // endless loop
    // }
    ```

4.  **`for...range` (iterating over collections like slices, arrays, maps, strings):**
    ```go
    // For slices/arrays
    nums := []int{2, 3, 4}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // For maps
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // For strings (iterates over Unicode code points/runes)
    for index, char_rune := range "Go" {
        fmt.Printf("%d: %c\n", index, char_rune)
    }
    ```

---

## Latihan

1. `Perhitungan Faktorial (Factorial Calculation)` Diberikan sebuah bilangan bulat non-negatif n, tulis algoritma untuk menghitung faktorialnya. Faktorial dari n (dilambangkan sebagai n!) adalah hasil perkalian semua bilangan bulat positif yang kurang dari atau sama dengan n.
- Faktorial didefinisikan sebagai:
   - `0! = 1`
   - `1! = 1`
   - `n! = n * (n-1) * (n-2) * ... * 1`
- Contoh:
   - Input: `5`
   - Output: `120` (karena 5! = 5 * 4 * 3 * 2 * 1)

2.  `Check Palindrome` Buat program untuk memeriksa apakah sebuah string adalah palindrome. Palindrome adalah kata, frasa, angka, atau urutan karakter lainnya yang dibaca sama dari depan maupun belakang.
- Contoh:
   - Input: `"radar"`
   - Output: `true` (karena "radar" dibaca sama dari depan dan belakang)
   - Input: `"hello"`
   - Output: `false` (karena "hello" tidak sama jika dibaca terbalik)

3. `Pencarian Bilangan Paling Besar` Buat program untuk mencari bilangan terbesar dalam sebuah slice (array dinamis) dari bilangan bulat.
- Contoh:
   - Input: `[3, 5, 2, 8, 1]`
   - Output: `8` (karena 8 adalah bilangan terbesar dalam slice tersebut)


Selamat belajar!