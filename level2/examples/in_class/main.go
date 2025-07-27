package main

import (
	"fmt"
	"strconv"
)

// shared to public, bisa di akses dari package lain
func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// private function, hanya bisa diakses dari package ini
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		//
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil // nil means no error
}

func divideAlternative(a, b int) (int, error) {
	if b == 0 {
		err := fmt.Errorf("cannot divide by zero")
		return 0, err
	}

	result := a / b

	return result, nil // nil means no error
}

// return multiple values
func dividePartTwo(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}

	result = a / b
	return
}

// function with slice return type
func getAllFirstNameEmployee() (fullname []string, err error) {
	return
}

// function with no return value
func doSomething() {

}

// function with bool return type
func isEvenNumber(number int) (res bool, err error) {
	return
}

// nil sama seperti null di bahasa lain, artinya tidak ada nilai

// pass by value
func swap(a, b int) (int, int) {
	return b, a
}

// pass by reference
func swapByReference(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	return
}

func getEmployeeName() (name []*string, err error) {

	return nil, nil
}

// other function params
func getName(name string, age int, isActive ...bool) (string, int, bool) {
	// isActive is a variadic parameter, it can accept zero or more than one boolean values
	return name, age, isActive[0] // if isActive is empty, this will cause a panic
}

func getName2(name string, age int, isActive bool) (string, int, bool) {
	return name, age, isActive
}

func returnError() (err error) {
	// this function will return an error
	err = fmt.Errorf("this is an error")
	return
}

func main() {
	maxValue := GetMax(10, 20)

	fmt.Println("nilai max :: ", maxValue)

	fmt.Println("nilai min :: ", getMin(10, 20))

	// java throws exception
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hasil bagi :: ", result)

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("hasil bagi :: ", result)
	}

	// convert from string to int
	intValue, err := strconv.Atoi("10abc")
	if err != nil {
		fmt.Println("error converting string to int:", err)
	}
	fmt.Println("nilai dari err :: ", err)
	fmt.Println("nilai intValue :: ", intValue)

	fmt.Println()
	// function with multiple return values part two
	result, err = dividePartTwo(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("hasil bagi part two :: ", result)
	}

	// example panic in golang
	// defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered from panic:", r)
	//	}
	//}()

	//var employeeName []string
	//// panic: runtime error: index out of range [1] with length 0
	//fmt.Println(employeeName[1])

	fmt.Println("\nThis line will not be executed due to panic above")

	// pass by value
	a := 10
	b := 20
	fmt.Println("get max value :: ", GetMax(a, b))

	a, b = swap(a, b)
	fmt.Printf("\nswap result :: a = %d || b = %d", a, b)

	// pass by reference
	c := 10
	d := 20
	swapByReference(&c, &d)
	fmt.Printf("\nswap by reference result :: c = %d || d = %d", c, d)

	// default value of pointer is nil
	var cekAngka *int
	var notPointerVar int64

	fmt.Println("\n\ncek default value :: ", cekAngka)
	fmt.Println("notPointerVar :: ", notPointerVar)

	if cekAngka == nil {
		fmt.Println("cek default value :: ", cekAngka)
	}

	if notPointerVar == 0 {
		fmt.Println("notPointerVar is zero")
	}

	tempAngka := 10
	cekAngka = &tempAngka

	fmt.Println("cekAngka is not nil, value is :: ", *cekAngka)
	if *cekAngka > 10 {
		fmt.Println("cekAngka is greater than 10")
	}

	tempAngka = 20
	fmt.Println("cekAngka is still not nil, value is :: ", *cekAngka)

	// example call new function
	//name, age := getName("Charles", 20, true, false, true)

	//name2, age2 := getName2("Charles", 20, true)

	if err = returnError(); err != nil {
		fmt.Println("error:", err)
	}
	// above is equivalent to:
	err = returnError()
	if err != nil {
		fmt.Println("error:", err)
	}

}
