package main

import "fmt"

func main() {

	/*
		Input: nums = [1,3,5,6], target = 5
		Output: 2
	*/

	var i []int = []int{1, 3, 5, 7, 9, 11}
	t := 6

	fmt.Println(searchInsert(i, t))

	// fmt.Println(0 / 2)
}

func searchInsert(nums []int, target int) int {
	var left int
	var right = len(nums) - 1
	var pivot int

	for left <= right {
		pivot = left + (right-left)/2

		if nums[pivot] > target {
			right = pivot - 1
		} else if nums[pivot] < target {
			left = pivot + 1
		} else if nums[pivot] == target {
			return pivot
		}
	}

	if nums[pivot] > target {
		return pivot
	} else {
		return pivot + 1
	}
}
