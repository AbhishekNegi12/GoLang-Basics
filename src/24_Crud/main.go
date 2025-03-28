package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
)

type Todo struct{
	userID int `json:"userId"`
	Id int `json:"id`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	fmt.Println("Learning Crud Operation in Golang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error :",err);
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		fmt.Println("Error in getting Response :",res.Status);
		return
	}

	// Generic method of reading response
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil{
	// 	fmt.Println("Error Reading:",err);
	// }
	// fmt.Println("Data :",string(data));

	var todo Todo
	err = json.NewDecoder(res.Body).Decode((&todo))
	if err !=nil{
		fmt.Println("Error Reader : ",err)
	}
	fmt.Println("Todo :",todo);

}
