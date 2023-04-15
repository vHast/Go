package main

import (
	"fmt"
)

// * Making decisions

func main() {

	// To get a Boolean value, you usually use a comparison operator

	num := 6 // int 
	condition := num % 2 == 0 // true

	fmt.Println(condition)

	fmt.Println(num % 2 == 0) // You can also check for Boolean values like this

	// * Logical Operators in GO

	numExtra := 6
	fmt.Println(numExtra == 0) // False

	

	var operatorNumOne int = 5
	conditionOne := operatorNumOne > 2 && operatorNumOne < 9

	fmt.Println(conditionOne)

}