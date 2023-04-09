package main

import "fmt"

func main() {
	var num1 = 5				// type inferred
	var num2 int				// *explicitly typed

	fmt.Println(num1)		// 5
	fmt.Println(num2)		// 
}