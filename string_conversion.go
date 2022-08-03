// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// the most comman string constant are Atio (string to int) and Itoa(int to string)
// 	i, _ := strconv.Atoi("-42")
// 	s := strconv.Itoa(-42)
// 	b, _ := strconv.ParseBool("true")
// 	f, _ := strconv.ParseFloat("3.1415", 64)
// 	j, _ := strconv.ParseInt("-42", 10, 64)

// 	fmt.Println(i, s)
// 	fmt.Println(b)
// 	fmt.Println(f)
// 	fmt.Println(j)

// }

// package main

// import "fmt"

// func main() {
// 	/* an int array with 5 elements */
// 	var balance = []int{1000, 2, 3, 17, 50}
// 	// var avg float32

// 	/* pass array as an argument */
// 	avge := getAverage(balance, len(balance))

// 	/* output the returned value */
// 	fmt.Println("Average value is:  ", avge)
// 	fmt.Println()
// }
// func getAverage(arr []int, size int) func() float32 {
// 	var i, sum int
// 	var avg float32

// 	for i = 0; i < size; i++ {
// 		sum += arr[i]
// 	}

// 	avg = float32(sum / size)
// 	return func() float32 {
// 		c := avg + 1
// 		return c
// 	}
// }

// package main

// /* define a circle */
// type Circle struct {
//     x,y,radius float64
// }

// /* define a method for circle */
// func(circle Circle) area() float64 {
//     return math.Pi * circle.radius * circle.radius
// }

// func main() {
//     circle := Circle {x:0, y:0, radius:5}
//               fmt.Printf("Circle area: %f", circle.area())
// }

package main

func main() {

	// a := RomanToInt("VIIIIII")

	// fmt.Println(a)

	Prime_number()
}

var RomanInts = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var Exceptions = map[string]int{
	// "IIII":  4,
	// "IIIII": 5,s
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func RomanToInt(s string) int {
	n := len(s)
	result := 0
	i := 0

	for ; i < n-1; i++ {
		if val, found := Exceptions[s[i:i+2]]; found {
			result += val
			i++
		} else {
			result += RomanInts[s[i]]
		}
	}

	if i == n-1 {
		result += RomanInts[s[n-1]]
	}

	return result
}
