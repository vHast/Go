package main

import "fmt"

func valueOne() bool {
	fmt.Println("valueOne gets called")
	return true
}

func valueTwo() bool {
	fmt.Println("valueTwo gets called")
	return false
}

func main() {

	if valueOne() && valueTwo() {
		fmt.Println("Shouldn't print")
	}

	if valueTwo() && valueOne() {
		fmt.Println("Shouldn't print")
	}

	if valueOne() || valueTwo() {
		fmt.Println("The boolean expression is true")
	}

	if valueTwo() || valueOne() {
		fmt.Println("The boolean expression is true")
	}
}
