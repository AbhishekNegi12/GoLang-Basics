package main

import "fmt"

func main() {
	// In Go, a map is a data structure that provides an unordered collection of
	// key-value pairs, where each key must be unique.
	fmt.Println("Learning Maps")

	//creating maps
	mp := make(map[string]int)

	//inseting key value pairs
	mp["Abhishek"] = 100
	mp["Diya"] = 85
	mp["Bhanu"] = 65
	mp["Siya"] = 88

	//Accessing the whole map
	fmt.Println(mp)

	//Accessing a particular value
	fmt.Println("Marks of Abhishek :", mp["Abhishek"])

	//deleting any key
	// delete(mp,"Siya")
	// fmt.Println(mp)

	// Modifying values
	mp["Siya"] = 95
	fmt.Println(mp)

	//checking if a key exist
	//returns 2 things
	//1.its value
	// 2. and a boolean value for whether this key exist or not
	x, exist := mp["David"]
	fmt.Println("David marks  :", x)
	// The existence of a key is checked using the second return value of a map lookup
	fmt.Println("David exist:", exist)

	//Iterating over a map
	for index, value := range mp {
		fmt.Printf("%s : %d\n", index, value)
	}

	//creating a map using literal
	person := map[string]int{
		"Abhishek": 100,
		"ozil":     90,
		"Alex":     85,
		"Bhanu":    66,
	}
	fmt.Println(person)

	for key, value := range person {
		fmt.Printf("%s : %d\n", key, value)
	}
}
