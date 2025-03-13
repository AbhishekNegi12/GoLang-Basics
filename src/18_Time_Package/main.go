package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Learning Time Package")

	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Printf("The type is : %T\n",currTime)

	//will not give output
	// formatted := currTime.Format("dd-mm-yyyy")
	// fmt.Println("Formatted Time :",formatted)

	//Monday should be present to get the exavt day sam ewith date no othe date is acceptable
	// formatted := currTime.Format("02-01-2006, Monday")
	// fmt.Println("Formatted Time :",formatted)

	formatted := currTime.Format("02-01-2006, Monday")
	fmt.Println("Formatted Time :",formatted)

	formatted = currTime.Format("02/01/2006, 15:04:05")
	fmt.Println("Formatted Time :",formatted)

	formatted = currTime.Format("02-01-2006, 3:04 PM")
	fmt.Println("Formatted Time :",formatted)

	formatted = currTime.Format("2006/01/02, 15:04:05")
	fmt.Println("Formatted Time :",formatted)

	go_date := "2006-01-02"
	date_str := "2025-03-13"
	new_date,_ := time.Parse(go_date,date_str)
	fmt.Println(new_date)
}