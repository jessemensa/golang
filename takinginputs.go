package main

import (
	"fmt"
	"strconv"
)

// package for converting strings to integers

func main() {
	fmt.Println("Enter your first name: ") // take input from user
	var firstName string
	fmt.Scanln(&firstName) // take input from user

	fmt.Println("Enter your last name: ")
	var lastName string
	fmt.Scanln(&lastName)

	fmt.Print("Your first name is: ")
	fmt.Println(firstName)

	fmt.Print("Your last name is: ")
	fmt.Println(lastName)

	// concatenation
	fmt.Println("Your name is: " + firstName + " " + lastName)

	// adding two numbers
	var firstNumber string
	var secondNumber string

	fmt.Println("Enter the first integer: ") // user prompt
	fmt.Scan(&firstNumber)                   // take input from user

	fmt.Println("Enter the second integer: ") // user prompt
	fmt.Scan(&secondNumber)                   // take input from user

	fmt.Println(firstNumber + secondNumber)

	// converting strings to integers
	var firstNumberAgain string
	var secondNumberAgain string

	var firstInt int
	var secondInt int

	fmt.Println("Enter the first integer: ") // user prompt
	fmt.Scan(&firstNumberAgain)              // take input from user

	fmt.Println("Enter the second integer: ") // user prompt
	fmt.Scan(&secondNumberAgain)              // take input from user

	firstInt, _ = strconv.Atoi(firstNumberAgain)   // convert string to integer
	secondInt, _ = strconv.Atoi(secondNumberAgain) // convert string to integer

	fmt.Println(firstInt + secondInt)       // print the sum of the two integers
	fmt.Println(firstNumber + secondNumber) // print the concatenation of the two strings

	// ASSIGNMENT OPERATIONS IN GO
	// ==, += -= , *= , /= %= , &= , |= , ^= , &=^ , <<= , >>=

	var a, b, c, d int = 100, 50, 25, 4

	a *= b
	b /= c
	c %= d
	fmt.Println("d =", d)

	// Boolean
	var myBool bool = true
	fmt.Println("myBool =", myBool)

	var anotherBool bool = false
	fmt.Println("anotherBool =", anotherBool)

}
