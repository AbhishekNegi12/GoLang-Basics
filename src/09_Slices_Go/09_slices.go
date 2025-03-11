package main

import "fmt"

func main(){
	fmt.Println("Learning Slices")

	// Declare and initialize a slice
	num :=[]int{1,2,3,4,5}
	fmt.Println(num)

	//Length of slice
	fmt.Println("Length of num is :",len(num))

	//knowing the datatype
	fmt.Printf("Datatype of num is : %T\n",num)

	//Append in slice
	num = append(num, 99,100,12)
	fmt.Println("Slice num after appending few elements :",num)

	fmt.Println("***************-------------------**************")
	fmt.Println("More on Slices")

	//declaring a slice without capacity
	sl := []int{3,6,9,12}
	
	//Printing slice, length, capacity
	fmt.Println("Slice :",sl)
	fmt.Println("Slice Length :",len(sl))
	fmt.Println("Slice Capacity :",cap(sl))

	// Creating a slice with make: length = 3, capacity = 5
	//Use of "make" function
	arr2 := make([]int, 3,5)
	fmt.Println("Slice :",arr2)
	fmt.Println("Slice Length :",len(arr2))
	fmt.Println("Slice Capacity :",cap(arr2))

	//decla;ring empty slice
	var num2 = []string{}
	fmt.Println("Var2 :",num2)
	num3 := make([]int,0)
	fmt.Println(num3)

	//caapcity double itself *2
	arr2 = append(arr2, 12,15)
	fmt.Println("Slice :",arr2)
	//Magic
	arr2 = append(arr2, 18,21)
	fmt.Println("Capacity is double :",arr2)
	fmt.Println("Slice Length :",len(arr2))
	fmt.Println("Slice Capacity :",cap(arr2))

}