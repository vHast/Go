package main

import "fmt"
import "reflect"

func main() {
	var num1 = 5				// type inferred
	var num2 int				// *explicitly typed
	var rates float32 = 4.5 // declared as float32 and initialized

	fmt.Println(num1,num2,rates);

	// Another way todeclare and initialize a variable is to use the short variable declaration (:=)

	firstName := "Wei-Meng"
	fmt.Println(firstName) // Will print Wei-Meng
	fmt.Println(reflect.TypeOf(firstName)) // Will print the type of firstName

	// Variables declared with the (:=) operator can't be used outside function bodies

	//* STRINGS

	address := "The White House\n1600 Pennsylvania Avenue NW\nWashington, DC 20500" // A string can contain special characters as \n

	fmt.Println(address)

	quotation := 	`"Anyone who has ever made
	a mistake has never tried
	anything new."
	-- Albert Einstein`

	fmt.Println(quotation)
}