package main

import "fmt"

func main() {

	// included shorthand declaration
	// created an array without defining its size

	array := [...]int{83, 432, 4, 323, 3, 3, 3, 23, 423}
	array2 := [...]int{83, 423, 4, 323, 3, 3, 3, 23, 432}

	fmt.Println("array element:")
	for i := 0; i < 5; i++ {
		fmt.Print(array[i], " ")

	}

	fmt.Println()

	// getting the size of the array
	fmt.Println("length of array is :", len(array))

	// printing array without using any loop
	fmt.Println("printing array without using any loop")

	fmt.Println(array)

	// creating an array without any existing array

	array3 := array

	array3[0] = 107

	fmt.Println("creating an array without any existing array")

	fmt.Println(array3)

	fmt.Println("Golang program to compare two arrays using equal to (==) operator")

	if array == array2 {

		fmt.Println("arr1 = arr2")

	}

	if array == array3 {
		fmt.Println("arr1 = arr3")
	} else {
		fmt.Println("arr1 != arr2")

	}

	fmt.Println("Golang program to calculate the sum of all array elements")

	var sum int
	for i := 0; i < array[5]; i++ {
		sum = sum + array[i]
	}

	fmt.Println(sum)

	fmt.Println("Golang program to find the largest elements from the array")
	fmt.Println(array)

	g := len(array)
	var D int

	for i := 0; i < g; i++ {
		if D < array[i] {

			D = array[i]
			// break
		}
	}
	fmt.Println(D)

	fmt.Println("Golang program to find the second largest elements from the array")

	// var E []int
	var first int
	var second int
	fmt.Println("g", g)

	// var check int
	for i := 0; i < g; i++ {
		if first < array[i] {
			// check = array[i]
			// E = append(E, array[i])
			// break

			second = first

			first = array[i]

		}

	}

	fmt.Println("first", first)

	fmt.Println("second", second)

	// fmt.Println("E", E)

}

// var result []int

// for i := 0; i <= array[5]; i++ {
// 	result = append(result, array...)
// }
