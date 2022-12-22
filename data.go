// Arrays in GOLANG
package main

import "fmt"

// NOTE: WHEN AN ARRAY IS DECLARED, THE TYPE AND SIZE CANNOT BE CHANGED LATER

func main() {
	// create an array of three integers
	var numbers [3]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	// display the array
	fmt.Println(numbers)
	// what is this ??
	fmt.Println(numbers[0], numbers[1], numbers[2])

	// DECLARING AND INITIALISING ARRAYS
	var theNumbers = [5]int{100, 2, 3, 7, 50}
	words := [4]string{"hello", "world", "this", "is"}

	fmt.Println(theNumbers)
	fmt.Println(words)

	// using a for loop to define an array
	var numbers2 [20]int
	for x := 0; x < 20; x++ {
		numbers2[x] = x + 1
	}
	fmt.Println(numbers2)

	// creating a multi dimensional array
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	fmt.Println(matrix[0][0])
	fmt.Println(matrix[1][1])

	// duplicating an array
	numbers_1 := [3]int{1, 2, 3}
	fmt.Println(numbers_1)
	numbers_2 := numbers_1
	fmt.Println(numbers_2)
	fmt.Println(numbers_1)

	// Comparing arrays
	matrix_1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix_2 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix_3 := [3][3]int{
		{9, 9, 9},
		{9, 9, 9},
		{9, 9, 9},
	}

	// compare matrix_1 to matrix_2
	if matrix_1 == matrix_2 {
		fmt.Println("Matrix 1 and Matrix 2 are equal")
	} else {
		fmt.Println("Matrix 1 and Matrix 2 are not equal")
	}

	// compare matrix_1 to matrix_3
	if matrix_1 == matrix_3 {
		fmt.Println("Matrix 1 and Matrix 3 are equal")
	} else {
		fmt.Println("Matrix 1 and Matrix 3 are not equal")
	}

	fmt.Println(matrix_1)
	fmt.Println(matrix_2)
	fmt.Println(matrix_3)

}
