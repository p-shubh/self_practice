package main

import "fmt"

// taking a function as integer
// type pointer as a parameter

func PTF(a, b *int) {

	// dereferensing
	*a = 748
	*b = 657
}

func main() {

	var x, y = 876, 293

	fmt.Println("the value of x before function call ", x)
	fmt.Println("the value of y before function call ", y)

	// taking a pointer variable
	// and assign the address
	// of x to it

	var Pa *int = &x
	var Pas *int = &y

	// Calling the function by passing
	// Passing pointer to function

	PTF(Pa, Pas)

	// here the value of x is been changed by the value in the function is having

	fmt.Println("The value of x after function call is ", x)
	fmt.Println("The value of y after function call is ", y)

}
