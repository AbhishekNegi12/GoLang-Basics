package main

import (
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
)

func main() {
	// fmt.Println("Learning about File Handling in go")

	// file,err := os.Create("abhi.txt")
	// if err != nil{
	// 	fmt.Println("Error while opening file:",err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("Successfully creaated file")

	// content := "Hello world by Abhishek Negi"
	// byte,errors := io.WriteString(file,content+"\n")
	// if errors != nil{
	// 	fmt.Println("Error while opening file:",err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("Successfully creaated file of byte :",byte)

	//Reading data from file buffer
	/*
	file, err := os.Open("abhi.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("Successfully creaated file")

	//creating a buffer to read the file content
	buffer := make([]byte, 1024)

	//Read the file content into the buffer
	for {
		n, err := file.Read((buffer))
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file:", err)
			return
		}
		// fmt.Println("Error while Reading file")

		//Process the read content
		fmt.Println(string(buffer[:n]))
	}
	*/

	//Read the entire file into a byte slice 
	//*******SHORTCUT METHOD*************

	//now it is depriciated
	// content,err := ioutil.ReadFile("abhi.txt")

	content,err := os.ReadFile("abhi.txt")
	if err !=nil{
		fmt.Println("Error while reading the file")
		return
	}
	fmt.Println(string(content))
	// Using ioutil.ReadFile is convenient for scenarios where you want to read the entire
	// content of a file into memory. However, keep in mind that it might not be suitable for very large
	// files as it loads the entire content into memory at once. For larger files, reading in chunks as
	// shown in the previous example may be more appropriate
}
