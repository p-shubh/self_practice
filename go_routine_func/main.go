package main

import (
	"fmt"
	"time"
)

// main function

func main() {

	fmt.Println("Welcome! to the main function")

	// creating an anounomus function

	go func() {
		fmt.Println("Go routine main fucntion")
	}()

	for i := 1; i < 5; i++ {

		time.Sleep(1 * time.Second)

		fmt.Println("GOodbye! to the main function")

	}
}
