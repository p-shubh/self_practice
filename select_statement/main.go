/*
	IN GOLANG SELECT STATEMENT IS JUST LIKE SWITCH STATEMENT
	BUT IN THE SELECT STATEMENT, CASE STATEMENT REFERS TO
	COMMUNICATION, i.e. SENT OR RECIEVE OPERATION ON THE CHANNEL.
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

type Auther struct {
	Name      string
	Branch    string
	Particals int
	Salary    int
}

// Method with a reciever
// of auther type
func (a Auther) show() {

	fmt.Println("Auther's name:", a.Name)
	println("branch name:", a.Branch)
}

func main() {

	/*
			select{
			case SendOrRecieve1 : //Statement
			case SendOrRecieve2 : //Statement
			case SendOrRecieve3 : //Statement
			.........
			default //statement
	}*/

	/*

		IMPORTANT POINT:-

		Select statement waits until the communication (send or recieve operation) is
		prepeared for some cases to being.
	*/

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	var value int = 2

	// switch statement without an optional statement and
	// expression

	switch {
	case value == 1:
		fmt.Println("Hello")

	case value == 2:
		fmt.Println("Bonjour")

	case value == 3:
		fmt.Println("Namaste")

	default:
		fmt.Println("Invalid")
	}

	var Values string = "five"

	// Switch statement without default statement
	// Multiple value in case statement

	switch Values {
	case "one":
		fmt.Println("#c")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}

	/*
							TYPE SWITCH

		Type switch is used when you want to compare types. In case
		contain the type which is going to compare with the present
		in which is going to compare with the type prsent in the switch
		expression .

		Syntax:-
		--------

		Switch optstatement; typeSwitchexxpression{
		case typelist 1: Statement..
		case typelist 2: statement..
		...
		default: statement..
		}

	*/

	var Vals interface{}
	switch q := Vals.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}

	fmt.Println()
	fmt.Println("Index value: ", strings.Index("shubham", "m"))
	fmt.Println("Time: ", time.Now().Unix())

	fmt.Println("welcome to main() function")

	defer mul(7832, 7928)

	mul(36, 782)

	defer mul(32, 3243)
	defer mul(234, 6532)
	mul(326, 7246)

	show()

	res := Auther{
		Name:   "Shubham",
		Branch: "bjhber",
	}

	res.show()
}

// DEFER KEYWORD IN GOLANG

// FUNCTION

// defer func(parameter_list Type)return_type{
// code
// }

// METHOD
// defer func (reciever type) method_name(parameter_list){
// code
// }

// defer func (parameter_list)(return_type){
// code
// }

// GO programing to illustrate the
// concept of init() function

// Declaration of the main package

func init() {
	fmt.Println("Welcome to init function")
}

func init() {
	fmt.Println("Hello! init function")
}

func mul(a1, a2 int) int {

	c, _ := fmt.Println("the sum is", a1+a2)

	return c
}

func show() {
	fmt.Println("Hellow! geek for geeks")
}

// 	wg.Add(2)
// 	go foo()
// 	go bar()
// 	wg.Wait()

// 	fmt.Println(Adder(698, 98979, 9678799, 7687989, 897))
// 	fmt.Println(Sqrt(49))

// }

// func foo() {
// 	for i := 0; i < 45; i++ {
// 		fmt.Println("Foo:", i)
// 	}
// 	wg.Done()
// }

// func bar() {
// 	for i := 0; i < 45; i++ {
// 		fmt.Println("Bar:", i)
// 	}
// 	wg.Done()
// }

// func Adder(xs ...int) int {
// 	res := 0
// 	for _, v := range xs {
// 		res += v
// 	}
// 	return res
// }

// func Sqrt(x float64) float64 {
// 	z := 0.0
// 	for i := 0; i < 1000; i++ {
// 		z -= (z*z - x) / (2 * x)
// 	}
// 	return z
// }
