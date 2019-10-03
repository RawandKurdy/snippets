package main

import "fmt"

// Array Max Consecutive Sum - Interview Question -
// Asked by Microsoft, LinkedIn, Amazon and Samsung
// Maximum possible sum you can get from one of contiguous subarrays.
// In Golang
func arrayMaxConsecutiveSum2(inputArray []int) int {
	var sumMax int
	var sumSteps int
	var maxNumber int

	for _, number := range inputArray {
		sumSteps += number

		if maxNumber == 0 || maxNumber < number {
			maxNumber = number
		}

		if sumSteps < 0 {
			sumSteps = 0
		}
		if sumMax < sumSteps {
			sumMax = sumSteps
		}
	}

	if sumMax != 0 {
		return sumMax
	}
	return maxNumber
}

func main() {
	// Example
	arr := []int{1, -2, 3, -4, 5, -3, 2, 2, 2}
	fmt.Println(arrayMaxConsecutiveSum2(arr)) // 8
}
