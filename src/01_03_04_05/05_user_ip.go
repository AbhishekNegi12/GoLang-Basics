package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name : ")

	// var name string
	// fmt.Scan(&name)

	// fmt.Printf("My name is : %s",name)

	//using bufio package
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	// fmt.Printf("Hello Mr. %s", name)
	fmt.Println("Hello Mr.", name)
}
