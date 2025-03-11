package main

import "fmt"

func main() {
	fmt.Println("Learning Switch-Cae in golnag")

	// Example 1: Basic switch statement
	var ch int
	fmt.Print("Enter your choice : ")
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Enter valid choice for knowing the day")
	}

	// Example 2: Switch statement with multiple values
	var choice string
	fmt.Print("Enter your choice of month : ")
	fmt.Scan(&choice)

	switch choice {
	case "December", "January", "February", "March":
		fmt.Println("Season is winter")
	case "April", "May", "june":
		fmt.Println("Season is Spring")
	default:
		fmt.Println("Other Season")
	}

	// Example 3: Switch with expression
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
