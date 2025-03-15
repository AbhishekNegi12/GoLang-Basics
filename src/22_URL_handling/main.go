package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning 'URL'")

	myurl := "https://leetcode.com/problems/add-two-numbers/"
	fmt.Printf("Type of url : %T\n", myurl)

// 	The url.Parse function is used to parse a string into a URL object
//  (url.URL struct). This allows you to break down a URL into its individual components,
//  such as scheme, host, path, and query parameters.
	parseURL, err := url.Parse(myurl)
	if err!=nil{
		fmt.Println("We can't parse URL")
		return
	}
	fmt.Printf("Type of URL : %T\n",parseURL)

	// Accessing URL Components:
	fmt.Println("Scheme :",parseURL.Scheme)
	fmt.Println("Host :",parseURL.Host)
	fmt.Println("Host :",parseURL.Path)
	fmt.Println("Raw Query :",parseURL.RawQuery)

	// String(): Converts a URL object back to its string representation
	newUrl := parseURL.String();
	fmt.Println("URL :",newUrl);
}
