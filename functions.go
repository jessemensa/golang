package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

// functions with no return values
func DisplayUpper(x string) {
	fmt.Println("Original text: ", x)
	fmt.Println("Revised text: ", strings.ToUpper(x))
}

// functions with multiple return values
func rectStuff(length int, width int) (int, int) {
	a := length * width
	c := length + length + width + width
	return a, c
}

// functions with different types can be returned
func circleStuff(radius int) (int, float32) {
	d := radius * 2
	c := 2 * 3.14 * float32(radius)
	return d, c
}

func circleStuffOne(radius int) (d int, c float32) {
	d = radius * 2
	c = 2 * math.Pi * float32(radius)
	return
}

// VARIADIC FUNCTIONS
func sumN(numbers ...int) {
	sum := 0
	for i, num := range numbers {
		fmt.Println("Current element: ", num, "; Current index: ", i)
		sum += num
		fmt.Println("Sum of values: ", sum)
	}
}

// EMBEDDED FUNCTION
func passGenerator() func() string {
	length := 10
	return func() string {
		pwd := ""
		for i := 0; i < length; i++ {
			randomChar := string(rand.Intn(256))
			pwd += randomChar
		}
		return pwd
	}
}

func main() {
	fmt.Println("add function results: ", add(4, 6))

	a := "elizabeth"
	// call the function
	DisplayUpper(a)

	area, parameter := rectStuff(3, 5)
	fmt.Println("Area: ", area)
	fmt.Println("Parameter: ", parameter)

	diameter, circumference := circleStuff(5)
	fmt.Println("diameter: ", diameter)
	fmt.Println("circumference: ", circumference)

	diameter, circumference = circleStuffOne(5)
	fmt.Println("diameter: ", diameter)
	fmt.Println("circumference: ", circumference)

	// CLOSURES => FUNCTION THAT DOES NOT HAVE A NAME
	a1 := 4
	b1 := 5

	add := func() int {
		return a1 + b1
	}

	fmt.Println("Sum of a1 and b1: ", add())

	passGen := passGenerator()
	fmt.Println(passGen())
	fmt.Println(passGen())
	fmt.Println(passGen())
}
