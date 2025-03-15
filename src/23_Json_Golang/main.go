package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:isAdult"`
}

func main() {
	fmt.Println("learning Json")

	p1 := Person{Name: "Abhishek", Age:20, IsAdult:true}
	// fmt.Println("Person data is :",p1)

	//convert person to json
	//encode or marshall
	jsonData,err :=json.Marshal(p1)
	if err!=nil{
		fmt.Println("Erro Marshalling",err)
	}
	fmt.Println("Person data is :",string(jsonData))

	//decode or unmarshal
	var p1_data Person
	err = json.Unmarshal(jsonData,&p1_data)
	if err!=nil{
		fmt.Println("Error Unmarshalling",err)
	}
	fmt.Println("person data is :",p1_data)
}