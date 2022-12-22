package main

import "fmt"

// create a struct
type account struct {
	accountNumber string
	balance       float64
}

// using a struct literal to initialise a struct
type anotherAccount struct {
	accountNumber string
	balance       float64
}

// key value pairs with structs
type keyValueAccount struct {
	accountNumber string
	balance       float64
}

// PASSING A STRUCT CREATED WITH NEW
type anotherNewAccount struct {
	accountNumber string
	balance       float64
}

func closeAccount(CurrentAccount *anotherNewAccount) {
	CurrentAccount.accountNumber = "CLOSED" + CurrentAccount.accountNumber
}

// POINTERS AND STRUCTS
type pointerAccount struct {
	accountNumber string
	balance       float64
}

// NESTED STRUCTS
// => Go supports nested structs, which allow you to create structs and then use those structs as
// data type fields for other structs.

type nestedStructAccount struct {
	accountNumber string
	balance       float64
	owner         entity
}

type entity struct {
	id      string
	address string
}

func main() {
	var a account

	a.balance = 140
	a.accountNumber = "1234567890"
	fmt.Println(a)

	// retrieving values from a struct
	fmt.Println("The account number is: ", a.accountNumber)
	fmt.Println("The balance is: ", a.balance)

	var b = anotherAccount{"1234567890", 140}
	fmt.Println(b)

	//
	a1 := keyValueAccount{balance: 140, accountNumber: "1234567890"}
	fmt.Println(a1)

	//
	var aNew = new(anotherNewAccount)
	aNew.accountNumber = "1234567890"
	aNew.balance = 140
	fmt.Println(aNew.accountNumber)
	closeAccount(aNew)
	fmt.Println(aNew.accountNumber)

	// POINTERS AND STRUCTS
	aPointer := pointerAccount{accountNumber: "1234567890", balance: 140}
	p := &aPointer

	(*p).balance = 100
	fmt.Println(a)

	p.balance = 320
	fmt.Println(a)

	// NESTED STRUCTS
	e := entity{"000-00-0000", "123 Main Street"}
	a2 := nestedStructAccount{}
	a2.accountNumber = "1234567890"
	a2.balance = 140
	// assign the entity struct as a value in the account struct
	a2.owner = e
	fmt.Println(a2)

}
