package main

import "fmt"

func main() {

	// var array []int
	array := []int{83, 423, 4, 323, 3, 3, 3, 23, 432}
	// var result []int

	// for i := 0; i <= array[5]; i++ {
	// 	result = append(result, array...)
	// }
	fmt.Println("array element:")
	for i := 0; i < 5; i++ {
		fmt.Print(array[i], " ")

	}

}
