package main

import "fmt"

func modifyVal(x* int){
	*x = *x *10;
}

func main() {
	fmt.Println("Learning pointers in Go Lang")

	//declaring pointer
	// var num int
	// num = 2
	num := 2
	fmt.Println("Value of num is :", num)

	// var ptr* int
	// ptr = &num

	// var ptr* int = &num
	ptr := &num
	fmt.Println("Value of ptr is :", ptr)
	fmt.Println("Value inside ptr is :", *ptr)

	// var ptr2* int
	// if ptr2==nil{
	// 	fmt.Println("Pointer is not assigned")
	// }

	val :=12;
	modifyVal(&val)
	fmt.Println("Value contains : ",val)
}
