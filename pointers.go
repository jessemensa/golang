// POINTERS => A VARIABLE THAT STORES THE MEMORY ADDRESS OF
// ANOTHER VARIABLE WHERE THE VALUE IS STORED RATHER THAN STORING THE VALUE ITSELF
package main

import (
	"fmt"
	"strings"
)

// USING POINTERS WITH FUNCTIONS
// passing a pointer to a function, you are able
// to access and use the value stored in the variable the pointer references.
func isUpper(x *string) bool {
	if strings.ToUpper(*x) == *x {
		return true
	}
	return false
}

// convert a string to uppercase
func toUpper(x *string) {
	*x = strings.ToUpper(*x)
}

func main() {
	var a int = 20
	var b *int // pointer variable b
	// initialise a pointer variable
	b = &a // b stores the memory address of a

	fmt.Println(a)
	fmt.Println(b)

	// doing it in one line
	var a1 int = 20   // a stores the value 20
	var b1 *int = &a1 // b1 stores the memory address of a variable a1
	fmt.Println(a1)
	fmt.Println(b1)

	var c float32 = 10.3
	var d *float32 = &c

	var e string = "My string"
	var f *string = &e

	var g uint = 42
	var h *uint = &g
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(h)
	fmt.Println(g)

	// accessing pointers stored value
	var myVar int = 20
	b2 := &myVar
	fmt.Println(b2)  // print the memory address
	fmt.Println(*b2) // print the stored value inside the memory address

	// nil pointer
	var b3 *int     // create a pointer variable b
	fmt.Println(b3) // print the memory address

	// using pointers to change variables
	var a4 int = 20
	b4 := &a4

	fmt.Print("The value stored in a: ")
	fmt.Println(a4)

	fmt.Print("The memory address: ")
	fmt.Println(b)

	fmt.Print("Value stored in a (via pointer): ")
	fmt.Println(*b4)

	*b4 = 30
	fmt.Print("New value of a: ")
	fmt.Println(a4)

	// Array of pointers
	numbers := []int{100, 1000, 10000}

	var ctr int
	for ctr = 0; ctr < len(numbers); ctr++ {
		fmt.Println(numbers[ctr])
	}

	// create an array of pointers
	var numbersptr [3]*int
	// assign a pointer to each value in the original array
	// and store them in the new array
	for ctr := 0; ctr < len(numbersptr); ctr++ {
		numbersptr[ctr] = &numbers[ctr]
	}
	fmt.Println(numbersptr)

	// print the values the pointers point to
	fmt.Println(*numbersptr[0])
	fmt.Println(*numbersptr[1])
	fmt.Println(*numbersptr[2])

	// CHANGING THE VALUES INSIDE AN ARRAY
	numbers1 := []int{100, 1000, 10000}
	var ctr1 int
	for ctr1 = 0; ctr < len(numbers1); ctr1++ {
		fmt.Println(numbers1[ctr1])
	}

	// create an array of pointers
	var numbersptr1 [3]*int
	for ctr1 = 0; ctr1 < len(numbersptr1); ctr1++ {
		numbersptr1[ctr1] = &numbers1[ctr1]
	}
	fmt.Println(numbersptr1)
	*numbersptr[0] = 200
	fmt.Println(numbers1)

	// USING POINTERS WITH FUNCTIONS
	var message string = "hello world"
	messageptr := &message
	// return true if string
	fmt.Println(isUpper(messageptr))

	var anotherMessage string = "hello world"
	messageptr5 := &anotherMessage
	toUpper(messageptr5)
	fmt.Println(anotherMessage)

}
