package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning about Web request")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error get request")
		return
	}
	defer response.Body.Close()

	fmt.Printf("Type of response : %T\n", response)
	// fmt.Println("response :", response)

	//read the response body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading response")
		return
	}
	// fmt.Println("Response : ", data)
	fmt.Println("Response : ", string(data))
}
