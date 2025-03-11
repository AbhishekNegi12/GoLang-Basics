package main

import "fmt"

func main(){
	fmt.Println("Learning Conditionals in Go")

	//if-else block
	x := 10
	if x>5{
		fmt.Println("x is greater than 5")
	}else{
		fmt.Println("x is smaller than 10")
	}

	//if-elseif block
	y := 10
	if x>10{
		fmt.Println("y is greater than 10")
	}else if y>5{
		fmt.Println("y is greater than 5 but smaller than 10")
	}else{
		fmt.Println("y is smaller than 5 and 10 both")
	}

	// Example 4: Checking both conditions using &&
	if x > 5 && y < 10 {
		fmt.Println("Both conditions are true")
	} else {
		fmt.Println("At least one condition is false")
	}
}
