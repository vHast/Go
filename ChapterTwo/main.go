package main

import (
	"fmt"
	"reflect"
	"time"
	"strconv"
)

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

	//* Performing TYPE CONVERSIONS

	start := time.Now() // What type of variable is start?

	fmt.Printf("%T\n", start)

	// %T is a GO syntax representation of the type of the value
	// \n is a line feed or newline
	// start is the target variable

	fmt.Println(reflect.TypeOf(start)) // time.Time
	fmt.Println(reflect.ValueOf(start).Kind()) // struct

	// * Converting a variable's type

	var newAge int // Declaration
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%d", &newAge) // Scanf will scan the input and transfer the value to newAge, %d stands for base 10 numbers which is a format specifier
	fmt.Println("You entered:", newAge)

	// If we enter a string it won't read anything

	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input) // %s stands for string, this will store a string value in input

	// After the input is read, you can use the strconv package's Atoi() function to convert the string into an integer value

	newAges, err := strconv.Atoi(input) // Convert string to int

	// The Atoi() function returns two values, the result of the conversion and the error

	// To check if the error ocurred during the conversion, check if the err variable contains a nil value

	if err != nil {
		fmt.Println(err) // An error ocurred
	} else {
		fmt.Println("Your age is", newAges)
	}

	// We can convert string values to specific types, such as Boolean, floating point numbers, integers, you may use the Parse functions

	b, bErr := strconv.ParseBool("t")
	fmt.Println(b) // true
	fmt.Println(bErr) // <nil>
	fmt.Printf("%T\n", bErr) // bool

	f, fErr := strconv.ParseFloat("3.1415", 64) //First value for ParseFloat is the value and then the specified bitsize for the float

	fmt.Println(f) // 3.1415
	fmt.Println(fErr) // <nil>
	fmt.Printf("%T\n", f) // float64

	i, iErr := strconv.ParseInt("-18.56", 10, 64) // First value for ParseInt is the value, then the base and the bitsize for that int

	fmt.Println(i) // 0
	fmt.Println(iErr)			//ERROR: Parsing "18.56": invalid syntax, should be a float
	fmt.Printf("%T\n", i)	//int64

	u1, u1Err := strconv.ParseUint("18", 10, 64)
	fmt.Println(u1) // 18
	fmt.Println(u1Err) // <nil>
	fmt.Printf("%T\n", u1) // uint64

	// To convert between the various numeric data types like int and float, you can simply use the int(), float32() and float64() functions

	number1 := 5
	number2 := float32(number1)
	number3 := float64(number2)
	number4 := float32(number3)
	number5 := int(number4)

	fmt.Printf("%T\n", number1) // int
	fmt.Printf("%T\n", number2) // float32
	fmt.Printf("%T\n", number3) // float64
	fmt.Printf("%T\n", number4) // float32
	fmt.Printf("%T\n", number5) // int 

	// * INTERPOLATING STRINGS

	queue := 5
	nameQueue := "John"

	// If you would like to concatenate you would do something like this 

	// * s := nameQueue + ", your queue number is:" + queue

	// This is wrong since queue is an integer and you CAN'T CONCATENATE string AND int VALUES, to fix this you may convert the int variable into a string

	// * s := nameQueue * ", your queue number is" + strconv.Itoa(queue)

	// This could work however it can be terrible when using a lot of variables

	// * A better solution would be to use Sprintf() function from the fmt package

	s := fmt.Sprintf("%s, your queue number is %d", nameQueue, queue)

	// The Sprintf() function formats a string baesd on the formatting verbs(such as %d and %s)

	fmt.Println(s);

}