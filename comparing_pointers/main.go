// go program to illustrate the
// concept of compating two pointers

package main

import "fmt"

func main() {

	val1 := 8943
	val2 := 645

	// creating and initializing the pointer

	// var P1 *int

	P1 := &val1

	P2 := &val2

	P3 := &val1

	// compairing each pointer with each other
	// using == operator

	res1 := P1 != P2
	fmt.Println(P1, P2)
	fmt.Println("Is P1 pointer is equal to P2 pointer res1", res1)

	res2 := &P1 != &P2
	fmt.Println(&P1, &P2)
	fmt.Println("Is P1 pointer is equal to P2 pointer res2", res2)

	res3 := P1 != P3
	fmt.Println(P1, P3)
	fmt.Println("Is P1 pointer is equal to P3 Pointer res3", res3)

	res4 := &P1 != &P3
	fmt.Println(&P1, &P3)

	fmt.Println("Is P1 pointer is equal to P3 pointer res4", res4)

	res5 := P2 != P3
	fmt.Println(P2, P3)

	fmt.Println("Is P2 pointer is equal to P3 pointer res5", res5)

	res6 := &P2 != &P3
	fmt.Println(&P2, &P3)

	fmt.Println("Is P2 pointer is equal to P3 pointer res6", res6)
}
