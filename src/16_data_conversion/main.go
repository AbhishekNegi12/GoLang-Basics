package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Learning Data Conversion")

	var num int = 42
	fmt.Println(num)
	fmt.Printf("Type of num is : %T\n",num)

	// num = num + 1.23 error in golang

	var data float64 = float64(num)
	fmt.Println(data)
	fmt.Printf("Type of data is : %T\n",data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println(str)
	fmt.Printf("Type of num is : %T\n",str)

	x := "1234"
	// int_str := strconv.Atoi(x)
	int_str,_ := strconv.Atoi(x)
	fmt.Println(int_str)
	fmt.Printf("Type of num is : %T\n",int_str)

	x = "1.64"
	float_str,_ := strconv.ParseFloat(x,64)
	fmt.Println(float_str)
	fmt.Printf("Type of num is : %T\n",float_str)
}
