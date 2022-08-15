/*
	IN GOLANG SELECT STATEMENT IS JUST LIKE SWITCH STATEMENT
	BUT IN THE SELECT STATEMENT, CASE STATEMENT REFERS TO
	COMMUNICATION, i.e. SENT OR RECIEVE OPERATION ON THE CHANNEL.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

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

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println(Adder(698, 98979, 9678799, 7687989, 897))
	fmt.Println(Sqrt(49))

}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}

func Adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
