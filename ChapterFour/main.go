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

}
