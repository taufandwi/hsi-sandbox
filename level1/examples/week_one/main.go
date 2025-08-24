package main

import (
	"fmt"
	"strconv"
)

func main() {
	//deklarasi variabel
	//var namaSantri string = "Taufan Dwi"
	//
	//var nik
	//
	//namaBuah := "Mangga"
	//angkaSaja := 10

	// data types and variable declaration
	var angka int
	var unsignedAngka uint64
	var angkaFloat float64
	var isTrue bool
	var nama string
	var angkaString string = "12345"

	angka = -5
	unsignedAngka = 50000
	angkaFloat = 3.14
	isTrue = false
	nama = "Taufan Dwi"

	fmt.Println(angka)
	fmt.Println(angkaFloat)
	fmt.Println(isTrue)
	fmt.Println(unsignedAngka)
	fmt.Println(nama)

	// konversi string ke int
	angka, _ = strconv.Atoi(angkaString)
	fmt.Println("konversi dari string ke int :: ", angka)

	// konversi int ke string
	nama = strconv.Itoa(angka)
	fmt.Println("konversi dari int ke string ::", nama)

	konv := fmt.Sprintf("konversi dari int ke string dengan sprintf :: %v\n || bool :: %v", angka, isTrue)
	fmt.Println(konv)

	// array and slice

	// deklarasi array
	//var numberOfArray [5]int64
	//var stringArray [3]string
	//var floatArray [4]float64

	// deklarasi slice
	//var numberOfSlice []int
	//var stringSlice []string
	//var floatSlice []float64

	// second option for slice and array declaration
	newNumberOfSlice := []int{12, 34, 56, 78, 90}
	newStringSlice := []string{"apple", "banana", "cherry"}

	//numberOfSlice = []int{1, 2, 3, 4, 5}
	//
	//stringArray = [3]string{"apple", "banana", "cherry"}

	fmt.Println("Array and Slice Example ::")
	fmt.Println(newNumberOfSlice)
	fmt.Println(newStringSlice)

	newNumberOfSlice = append(newNumberOfSlice, 123)
	newStringSlice = append(newStringSlice, "apple")

	fmt.Println("Array and Slice after append ::")
	fmt.Println(newNumberOfSlice)
	fmt.Println(newStringSlice)

	// map
	//var newMap map[string]interface{}{
	//	"nama":    "Taufan Dwi",
	//	"umur":    25,
	//	"alamat":  "Jl. Contoh No. 123",
	//	"isActive": true,
	//}

	// if else statement
	// jika
	// operators: ==, !=, <, >, <=, >=
	// logical operators: && all true, || salah satu, !
	angka = -10
	if angka <= 10 && angka >= -10 || angka == 0 {
		if unsignedAngka > 1000 {
			fmt.Println("Angka lebih besar dari 1000")
		} else {
			fmt.Println("Angka kurang dari atau sama dengan 1000")
		}
		fmt.Println("Angka lebih besar dari 0")

	} else if (angka < 0 && angka > -10) || angka == 0 {
		fmt.Println("Angka negatif")
	} else {
		fmt.Println("default case")
	}

	if !isTrue {
		fmt.Println("true")
	}

	// switch statement
	fmt.Println("switch statement example ::")

	angka = 15

	switch angka {
	case 5:
		fmt.Println("Angka adalah 5")
		break
	case 10:
		fmt.Println("Angka adalah 10")
		break
	default:
		fmt.Println("Angka tidak diketahui : ", angka)
	}

	// for loop
	fmt.Println("")
	fmt.Println("for loop example ::")

	// i := 0 inisialisasi variabel i
	// for (inisisi);(statement berjalannya sebuah perulangan):(increment) {}
	// i++ equal dengan i = i + 1
	// i + 1 , nilainya tidak masuk kedalam variabel i

	for i := 0; i < 20; i = i + 2 {
		fmt.Println("Perulangan ke-", i)
	}

	for i, j := 0, 10; i < 10 && j > 0; i, j = i+1, j-1 {
		fmt.Println("Perulangan ke-", i, "dan j adalah", j)
	}

	// for each loop
	for i, buah := range newStringSlice {
		fmt.Println("Buah ke-", i, "adalah", buah)
	}

	fmt.Println("")
	fmt.Println("second option")
	for i := 0; i < len(newStringSlice); i++ {
		fmt.Println("Buah ke-", i, "adalah", newStringSlice[i])
	}

	/*for {
		fmt.Println("This is an infinite loop. Press Ctrl+C to stop.")
		// You can add a break condition here if needed
		// For example, you could break after a certain number of iterations
		// or based on some other condition.
	}*/
}
