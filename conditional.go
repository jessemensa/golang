package main

import "fmt"

func main() {
	var age = 12

	if age > 16 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive")
	}

	// if statement can be used to check a false condition

	// declare a balance
	var accountBalance = 2000

	if (accountBalance - 500) < 0 {
		fmt.Println("You can't withdraw money")
	} else {
		fmt.Println("You can withdraw money")
	}

	if (accountBalance > 0) == false {
		fmt.Println("You can't withdraw money")
	} else {
		fmt.Println("You can withdraw money")
	}

	// Working with Multiple Conditions && and ||
	var userName string = "Jesse"
	var password string = "1234"

	if (userName == "Jesse") && (password == "1234") {
		fmt.Println("You are logged in")
	} else {
		fmt.Println("You are not logged in")
	}

	if (userName == "Jesse") || (password == "1234") {
		fmt.Println("You are logged in")
	} else {
		fmt.Println("You are not logged in")
	}

	// switch statement
	var color string = "red"

	switch color {
	case "Blue":
		fmt.Println("Color is blue")
	case "Red":
		fmt.Println("Color is red")
	case "Green":
		fmt.Println("Color is green")
	default:
		fmt.Println("Please choose a valid color.")
	}

	// fallthrough statement => when you want your program to flow from your current case into the next case.
	var number int = 4

	switch number {
	case 10:
		fmt.Println("Number is 10")
		number -= 1
		fallthrough
	case 9:
		fmt.Println("Number is 9")
		number -= 1
		fallthrough
	case 8:
		fmt.Println("Number is 8")
		number -= 1
	default:
		fmt.Println("Number is not 10, 9 or 8")
	}

}
