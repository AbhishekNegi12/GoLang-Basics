package main

import "fmt"

// Define a struct named Person
type Person struct {
	Fname string
	Lname string
	Age   int
}

type contact struct{
	email string
	phone string
}

type Address struct{
	house_no int
	area string
	state string
}

type Employee struct{
	e_details Person
	e_contact contact
	e_address Address
}

func main() {
	fmt.Println("Learning struct")

	// Create an instance of the Person struct
	var p1 Person
	fmt.Println("Abhishek Person :", p1)

	p1.Fname = "Abhishek"
	p1.Lname = "Negi"
	p1.Age = 20
	fmt.Println("Abhishek Person :", p1)

	// Access the fields of the struct
	fmt.Println("First Name:", p1.Fname)
	fmt.Println("Last Name:", p1.Lname)
	fmt.Println("Age:", p1.Age)

	// Method 2: Using a struct literal
	p2 := Person{
		Fname: "Bob",
		Lname: "Johnson",
		Age:   35,
	}
	// Access the fields of the struct
	fmt.Println("First Name:", p2.Fname)
	fmt.Println("Last Name:", p2.Lname)
	fmt.Println("Age:", p2.Age)

	// Method 3: Using the new keyword (returns a pointer to the struct)
	p3 := new(Person)
	p3.Fname = "Alisha"
	p3.Lname = "Roy"
	p3.Age = 21
	// Access the fields of the struct
	fmt.Println("First Name:", p3.Fname)
	fmt.Println("Last Name:", p3.Lname)
	fmt.Println("Age:", p3.Age)


	// Corrected Employee struct initialization
	emp1 := Employee{
		e_details : Person{
			Fname: "Abhishek",
			Lname: "Negi",
			Age: 20,
		},
		e_contact: contact{
			email: "44.negi@gmail.com",
			phone: "123456789",
		},
		e_address: Address{
			house_no: 202,
			area: "New Delhi",
			state: "Delhi",
		},
	}
	fmt.Println(emp1)

}
