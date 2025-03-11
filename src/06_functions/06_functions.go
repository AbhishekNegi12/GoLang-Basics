package main

import "fmt"

// Correct way
//func example() {
	// Function body
//}

// Incorrect way
// func example()
// {
	// Function body
// }

// Function without parameters and return type
func f1() {
	fmt.Println("Simple Function")
}

// Function with named return variable
func add(a, b int) int {
	return a + b
}

// no need to give variable in return
// Function with named return variable
func product(a int, b int) (result int) {
	result = a + b
	return
}

// Main function (entry point of the program)
func main() {
	fmt.Println("We are Learning fxns in go")
	f1()

	ans := add(5, 90)
	println("The sum is :", ans)

	res := product(5, 90)
	println("The product is :", res)
}
