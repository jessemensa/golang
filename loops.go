package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int = 0

	for ctr := 1; ctr <= 10; ctr++ {
		fmt.Println(ctr)
	}

	// Membershship inside loops
	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			fmt.Println(strconv.Itoa(a) + " is an even number")
		}
	}

	for a < 10 {
		if a%2 == 0 {
			fmt.Println(strconv.Itoa(a) + " is an even number")
		}
		a++ // increment manually
	}

	// GO WHILE LOOPS => GO DOES NOT HAVE A WHILE LOOP
	// something close to while loop is by removing the semicolons
	for a < 10 {
		if a%2 == 0 {
			fmt.Println(strconv.Itoa(a) + " is an even number")
		}
		a++ // increment manually
	}

	// Looping through a string
	var message string = "HELLO WORLD"
	fmt.Println(message)

	for idx := 0; idx < len(message); idx++ {
		fmt.Println(string(message[idx]))
	}

	// Range Function
	for idx, c := range message {
		fmt.Println(idx)       //index
		fmt.Println(string(c)) //value
	}

	// LOOP CONTROL STATEMENTS
	// break, continue & goto statements
	// goto => goto keyword sends the pro- gram flow to a different location identified by a label.
	var power2 int64 = 1
	var a1 int64 = 1

	for {
		if a1 >= 10 {
			break
		}

		fmt.Println("2 to the power of " + strconv.FormatInt(a1, 10) + " is equal to " + strconv.FormatInt(power2, 10))
		power2 += 2
		a1++
	}

	for ctr := 0; ctr < 10; ctr++ {
		if ctr%2 == 0 {
			continue
		}
		fmt.Println(strconv.Itoa(ctr) + " is an odd number")
	}

	var a2 int = 20
	var b2 int = 30
	fmt.Println("a = " + strconv.Itoa(a2))
	fmt.Println("b = " + strconv.Itoa(b2))

	if a2 > b2 {
		goto MESSAGE1 // this will jump the execution to where MESSAGE1
	} else {
		goto MESSAGE2
	}

MESSAGE1:
	fmt.Println("a is greater than b")
MESSAGE2:
	fmt.Println("b is greater than a")

}
