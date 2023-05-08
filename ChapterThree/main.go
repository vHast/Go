package main

import (
	"fmt"
)

// * Making decisions

func main() {

	// To get a Boolean value, you usually use a comparison operator

	num := 6                // int
	condition := num%2 == 0 // true

	fmt.Println(condition)

	fmt.Println(num%2 == 0) // You can also check for Boolean values like this

	// * Logical Operators in GO

	fmt.Println(num == 0) // false

	conditionOne := num > 2 && num < 9 // This checks if a number is more than 2 and less of a 9

	fmt.Println(conditionOne) // true

	conditionTwo := num < 2 || num > 9 // This checks if a number is less than 2 and more than a 9
	fmt.Println(conditionTwo)          // false

	conditionThree := num > 9 || num < 2 // This checks if a number is higher  than 9 or is a number inferior than 2

	fmt.Println(conditionThree) // false

	conditionFour := !(num > 9 || num < 2)
	fmt.Println(conditionFour) // true

}
