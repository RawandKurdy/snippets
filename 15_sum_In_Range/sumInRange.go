package main

import "fmt"

// Sum In Range - Interview Question - Asked by Palantir Technologies
// In Golang
// In O(n) time
func sumInRange(nums []int, queries [][]int) int {
	var totalSum int
	var prefixSums = map[int]int{}
	modulo := 1000000000 + 7

	// calculates the prefix-sums
	for index, num := range nums {
		if index == 0 {
			prefixSums[index] = num % modulo
		} else {
			prefixSums[index] = (prefixSums[index-1] + num) % modulo
		}
	}

	// Now adds up each separate range
	for _, query := range queries {
		fromIndex := query[0]
		toIndex := query[1]
		totalSum += prefixSums[toIndex]

		if fromIndex != 0 {
			totalSum -= prefixSums[fromIndex-1]
		}
		totalSum %= modulo
	}
	// if totalSum is negative we just add the modulo
	if totalSum < 0 {
		totalSum += modulo
	}
	return totalSum
}

func main() {
	// Example
	nums := []int{34, 19, 21, 5, 1, 10, 26, 46, 33, 10}
	queries := [][]int{{3, 7},
		{3, 4},
		{3, 7},
		{4, 5},
		{0, 5}}

	fmt.Println(sumInRange(nums, queries)) // 283
}
