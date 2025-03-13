package main
import "fmt"

func add(a,b int)int{
	return a+b;
}

func main(){

	//1.
	// fmt.Println("Learning about differ keyword")
	// fmt.Println("Program Started")
	// fmt.Println("Program Ended")
	// op:->
	// Learning about differ keyword
	// Program Started
	// Program Ended

	//2.
	// fmt.Println("Learning about differ keyword")
	// defer fmt.Println("Program Started")
	// fmt.Println("Program Ended")

	//op:->
	// Learning about differ keyword
	// Program Started
	// Program Ended

	//if two defer
	//will go into stack stage 
	//LIFO behaviour
	fmt.Println("Learning about differ keyword")
	res := add(45,652)
	defer fmt.Println("Sum is :",res)
	defer fmt.Println("Program Started")
	fmt.Println("Program Ended")

	//op:->
// 	Learning about differ keyword
// 	Program Ended
// 	Program Started
// 	Sum is : 697

}