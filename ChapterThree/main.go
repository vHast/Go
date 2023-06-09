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

	// * Using the if/else statement

	// Using logical and comparison operators, to generate a Boolean value so you can make decisions

	ifNum := 5
	ifCondition := ifNum%2 == 1 // true

	if ifCondition {
		fmt.Println("Number is odd")
	}

	// We can also check it explicitly

	if ifCondition == true {
		fmt.Println("Number is odd")
	}

	// You can also put a logical expression directly if the conditional part of the if statement

	if num%2 == 1 {
		fmt.Println("Number is odd")
	}

	// Parentheses around conditions are not required, unlike in Javascript

	if num%2 == 1 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	// * Short circuiting

	// Go evaluates conditions using a method known as short-curcuiting, as long as its raining or snowing you have to stay indoors

	if raining() || snowing() {
		fmt.Println("Stay indoors!")
	}

	// * Using IF with an initialization statement to keep the scope of your variables tight

	// When using the if statement, you can combine intialization together with the conditon statement, like this

	// This function returns an integer value together with a Boolean value, to indicate if there is an error code...

	v, err := doSomething()

	if err {
		fmt.Println("Error") // handle the error
	} else {
		fmt.Println(v)
	}

	// If the function has an error, you handle the error, otherwise, you can go ahead and use the value returned by the function

	if v, err := doSomething(); !err {
		fmt.Println(v)
	} else {
		fmt.Println("Error") // handle the error
	}

	// * Switch statements

	numSwitch := 5
	dayOfWeek := ""

	switch numSwitch {

	case 1:
		dayOfWeek = "Monday"
	case 2:
		dayOfWeek = "Tuesday"
	case 3:
		dayOfWeek = "Wednesday"
	case 4:
		dayOfWeek = "Thursday"
	case 5:
		dayOfWeek = "Friday"
	case 6:
		dayOfWeek = "Saturday"
	case 7:
		dayOfWeek = "Sunday"
	default:
		dayOfWeek = "--error--"
	}

	fmt.Println(dayOfWeek) // Friday

	// * Switch fallthroughs
	// When a fallthrough keyword is present in the switch statement, the next case is evaluated instead of transferring control out of the switch statement

	grade := "C"

	switch grade {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Absent")
	}

	// * Matching multiple cases

	switch grade {
	case "A", "B", "C", "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Undefined")
	}

	// * Switching without conditions
	// This construct makes it easy to perform conditional checks with the case expressions

	score := 79
	gradeCondition := ""

	switch {
	case score < 50:
		gradeCondition = "F"
	case score < 60:
		gradeCondition = "D"
	case score < 70:
		gradeCondition = "C"
	case score < 80:
		gradeCondition = "B"
	default:
		grade = "A"
	}
	fmt.Println(gradeCondition) // B

}

func raining() bool {
	fmt.Println("Check if it is raining now...")
	return true
}

func snowing() bool {
	fmt.Println("Check if it is snowing now...")
	return true
}

func doSomething() (int, bool) {
	return 5, false
}
