package main

import (
	"encoding/json"
	"fmt"
)

// simple struct definition
type Person struct {
	Name    string
	Age     int
	Address Address
}

// adding other struct fields
type Address struct {
	Street string
	City   string
}

// methods on structs
type Employee struct {
	Name     string `json:"name"`
	JobDesk  string `json:"job_desk"`
	IsActive bool   `json:"-"`
}

func (e *Employee) Greet() {
	fmt.Printf("Hello, my name is %s and I work as a %s.\n", e.Name, e.JobDesk)
}

func (e *Employee) GetJobDesk() string {
	return e.JobDesk
}

func (e *Employee) SetJobDesk(jobDesk string) {
	e.JobDesk = jobDesk
}

func main() {
	//p1 := Person{"charles", 20}

	p2 := Person{Name: "walker"}

	// print structs
	//fmt.Println(p1.Name, p1.Age)

	fmt.Printf("%v\n", p2)

	// init struct part 2
	address := Address{
		Street: "123 Main St",
		City:   "Springfield",
	}

	p1 := Person{
		Name:    "Charles",
		Age:     20,
		Address: address,
	}

	fmt.Printf("%v\n", p1)

	//another way to init struct
	var p4 Person

	p4.Name = "John"
	p4.Age = 18
	p4.Address = Address{
		Street: "Jalan Kemayoran",
		City:   "Jakarta",
	}

	fmt.Printf("\n%v\n", p4)

	p4.Address.City = "Bandung" // updating a field in the nested struct

	fmt.Printf("%v\n", p4)

	// using a pointer to a struct
	// var angka *int

	var p5 *Person
	p5 = &p4 // p5 now points to p4

	fmt.Printf("\nperson 5: %v\n", p5)

	p4.Address.Street = "Jalan Sudirman" // updating a field in the nested struct through pointer

	fmt.Printf("Updated person 5: %v\n", p5)

	// using methods on structs
	employee := Employee{Name: "Alice", JobDesk: "Software Engineer"}

	fmt.Printf("\n our employee :: %v\n", employee)

	employee.Greet()                                    // calling the Greet method
	fmt.Printf("Job Desk: %s\n", employee.GetJobDesk()) // calling the GetJobDesk method

	employee.SetJobDesk("Lead Software Engineer") // setting a new job desk
	// method diatas similar dengan employee.JobDesk = "Lead Software Engineer"

	fmt.Printf("Updated Job Desk: %s\n", employee.GetJobDesk()) // calling the GetJobDesk method again

	// print with json tag/marshal
	jsonData, err := json.Marshal(employee)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Printf("employee print with json :: %s\n", jsonData)

	// declare struct within function
	var order struct {
		Item  string  `json:"item"`
		Price float64 `json:"price"`
	}

	order.Item = "book"
	order.Price = 12.99999

	fmt.Printf("\nOrder Item: %s, Price: %.2f\n", order.Item, order.Price)
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Order JSON: %s\n", jsonOrder)

	// print struct with json tag part 2
	order2 := Order{
		OrderID:    "12345",
		CustomerID: "67890",
	}

	jsonOrder2, err := json.Marshal(order2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Order2 JSON: %s\n", jsonOrder2)

}

// example struct with json tag part 2
type Order struct {
	OrderID    string `gorm:"primaryKey" json:"order_id"`
	CustomerID string `bson:"customer_id"`
}
