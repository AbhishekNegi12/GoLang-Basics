package main

import "fmt"

func main() {
	var name string = "Abhishek Negi"
	fmt.Println(name)

	var version = 89
	fmt.Println(version)

	var x int = 12
	fmt.Println(x)

	var y float32 = 89.2000
	fmt.Println(y)

	// var flag bool = true
	var flag = true
	fmt.Println(flag)
	fmt.Printf("Type of flag : %T\n", flag)

	const pi = 3.14
	fmt.Println(pi)
	// will throw error
	// pi = 3.22

	//Shortcut of declaring variable
	e := 2.718
	fmt.Println(e)

}
