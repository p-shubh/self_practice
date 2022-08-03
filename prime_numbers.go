package main

import "fmt"

func Prime_number() {

	// count := 0
	var count int
	for i := 1; i < 100; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {

				count++

			}
		}
		if count == 2 {
			fmt.Println("i", i)
		}
		count = 0

	}

}
