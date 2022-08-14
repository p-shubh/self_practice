package main

import (
	"fmt"
	"time"
)

func Display(str string) {
	for i := 0; i < 6; i++ {

		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
	// return str
}

func main() {

	// calling go routine
	go Display("shubham Prajapati")

	// calling normal function
	Display("Geek for geeks")
}
