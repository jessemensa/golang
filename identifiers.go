// Identifiers

package main

import "fmt"

// global variable => can be accessed anywhere in the program
var globalMessage string = "hello world"

func main() {
	var Age int = 10
	var age_2 int = 20
	var _age20 int = 30
	var name string = "Jesse"
	var surname string = "Mensah"

	// declare two strings
	var message, email string
	message = "Hello Jesse"
	email = "jessemensah21@gmail.com"

	// dynamically typed variables
	anotherEmail := "jessemensah21@gmail.com"

	// local variable scope
	var anotherMessage string = "Hello World"

	fmt.Println(Age)
	fmt.Println(age_2)
	fmt.Println(_age20)
	fmt.Println(name, " ", surname)
	fmt.Println(message)
	fmt.Println(email)
	fmt.Println(anotherEmail)
	fmt.Println(anotherMessage)
}
