package main

import "fmt"

// can't handle float values
func divide1(a, b int) int {
	return a / b
}

// can't handle float values because input type is "int"
//Error
// func divide2(a, b int) float64 {
// 	return a / b
// }


func divide3(a, b float64) float64 {
	return a / b
}

//Error Handling
func Divide(a,b float64)(float64, error){
	if b==0 {
		return 0, fmt.Errorf("dominator must not be zero")
	}
	return a/b,nil
}

func main() {
	fmt.Println("Error Handling")

	ans1 := divide1(10, 4)
	fmt.Println("Result is :", ans1)

	// ans2 := divide2(10, 2)
	// fmt.Println("Result is :", ans2)

	ans3 := divide3(10, 0)
	fmt.Println("Result is :", ans3)

	//Instead using this use of "underscore"
	res,_ := Divide(10,0);
	fmt.Println("The result is :",res);

	//err declared but not used
	// res,err:= Divide(10,0);

	//we can use it
	res,err:= Divide(10,0);
	if err!=nil{
		fmt.Println("Error Handling");
	}
	fmt.Println("The divison is :",res);
}
