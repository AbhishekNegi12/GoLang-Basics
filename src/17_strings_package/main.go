package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Learning String package")

	data := "apple,orange,banana"
	// func strings.Split(s string, sep string) []string
	parts := strings.Split(data,",")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str,"two")
	fmt.Println(count)

	str = "     hello, Go!   "
	// func strings.TrimSpace(s string) string
	// TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
	trim := strings.TrimSpace(str)
	fmt.Println(trim)

	str1 := "Abhishek"
	str2 := "Negi"
	// res := strings.Join([]string{str1,str2},",")
	res := strings.Join([]string{str1,str2},"+")
	fmt.Println(res)
}
