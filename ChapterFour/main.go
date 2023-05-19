package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// If you want to print it in reverse...

	for i := 4; i >= 0; i-- {
		fmt.Println(i)
	}

	// Although a for loop technically has three sections (INIT, CONDITION, POST) the INIT and POST statements are actually optional

	max := 100
	a, b := 0, 1

	for b <= max {
		println(b)
		a, b = b, a+b
	}

	// In the previous code snippet, you only have the conditional expression in the for loop, you just need to check if b is less than or equal to max in order for the iteration to continue.
	// When b exceeds max, the loop stops

	// * Infinite loops
	// A foor loop without the three parts is an infinite loop

	// The following code waits for the user to input the string QUIT

	for {

		fmt.Println("Enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string:")
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}

	// ToUpper() converts the user input to uppercase before the comparison
	// This way the user can enter the QUIT string in any cases and the comparison will still work correctly

	// * Continue keyword

	for contN := 1; contN < 10; contN++ {
		if contN%2 == 0 {
			continue
		}
		fmt.Println(contN)
	}

	// * Iterating over a Range of Values
	// Iterating through arrays/slices

	// OS is an array of three elements

	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Android"
	OS[2] = "Window"

	// To iterate through each of the elements in the array, you use the for-range loop

	for i, v := range OS {
		fmt.Println(i, v)
	}

	// The range keyword returns the following values

	// i is the INDEX OF THE VALUE YOU'RE ACCESSING - in this case, the value is the OS array
	// v Each of the values in the OS array

}
