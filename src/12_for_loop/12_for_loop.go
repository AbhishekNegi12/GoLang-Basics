package main

import "fmt"

func main() {
	fmt.Println("Learninf for loop")
	// Example 1: Basic for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Printf("\n")

	// Example 2: Infinite loop with break statement
	cnt := 0
	for {
		fmt.Println("Infinite loop")
		cnt++
		if cnt == 10 {
			fmt.Println("Value is 10 now exiting loop")
			break
		}
	}

	//range keyword simplify looping over element
	// 1.over slice
	num := []int{1, 2, 3, 4, 5}
	for index,value := range num {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 2. over string
	data := "Hello world"
	for index,char := range data{
		fmt.Printf("Index : %d Value : %c\n",index,char)
	}
}
