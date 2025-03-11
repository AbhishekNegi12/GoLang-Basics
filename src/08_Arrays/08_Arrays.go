package main

import "fmt"

func main() {
	fmt.Println("Learning Arrays")

	// Declare an array of integers
	var arr[5]int
	//will store default value of int which is 0
	fmt.Println("Array :",arr)

	var str[10]string
	fmt.Println("strng is :",str)
	// %q is quouted
	fmt.Printf("str is :%q\n",str)

	// Declare and initialize an array to store a list of names
	var name [5]string
	name[0] = "Abhishek"
	name[1] = "Vivek"
	name[2] = "Bhanu"
	name[3] = "Diya"
	// Access and print array elements
	fmt.Println(name)

	// Initialize and assign values to an array in a single line
	numbers := [5]int{1,2,3,4,5}
	fmt.Println(numbers)

	//length of array
	fmt.Println("Lenght of array numbers is :",len(numbers))

	//Accessing any particular value
	fmt.Println(numbers[2])	
}
