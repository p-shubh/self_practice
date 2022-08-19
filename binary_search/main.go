package main

import "fmt"

func main() {

	/*
		Input: nums = [1,3,5,6], target = 5
		Output: 2
	*/

	var i []int = []int{1, 3, 5, 7}
	t := 6

	fmt.Println(searchInsert(i, t))
}

func searchInsert(k []int, t int) int {
	var result int

	c := len(k)

	s := c / 2 //2
	if s < t { //2<=5
		for i := s; i < c; i++ {
			if k[i] == t {
				result = i
			} else if s >= t {
				for i := 0; i <= s; i++ {
					if k[i] == t {
						result = i

					}

				}

			}

		}
	}

	if result == 0 {
		for j := 0; j <= c-1; j++ {
			if t > k[j] {
				result++
			}
		}
	}

	return result

}
