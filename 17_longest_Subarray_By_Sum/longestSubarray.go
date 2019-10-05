package main

import "fmt"

// Longest Subarray By Sum in an Array - Interview Question -
// Asked by Facebook & Palantir Technologies
// In Golang
func findLongestSubarrayBySum(sum int, arr []int) []int {
	var prefixSumsToIndex = map[int]int{}
	var fromIndex = -1
	var toIndex = -1
	var sumExistAt = -1

	// Calculates the prefixSums
	var previousSum = 0
	for index, number := range arr {
		var newSum = previousSum + number
		previousSum = newSum
		prefixSumsToIndex[newSum] = index + 1
		if number == sum {
			sumExistAt = index + 1
		}
	}

	// Sees if such sum exists within the array
	// initializes the indices with initial range if exists
	if index, ok := prefixSumsToIndex[sum]; ok {
		fromIndex = 1
		toIndex = index
	}

	// Searches the array for a wider range if available
	for numSum, indexofSum := range prefixSumsToIndex {
		indexOfSumWithS, ok := prefixSumsToIndex[numSum+sum]
		if ok {
			tmpFromIndex := indexofSum + 1
			tmpToIndex := indexOfSumWithS
			if (toIndex - fromIndex) < (tmpToIndex - tmpFromIndex) {
				toIndex = tmpToIndex
				fromIndex = tmpFromIndex
			}
		}
	}

	// Returns depending on if indices has been found or not
	if fromIndex != -1 && toIndex != -1 {
		return []int{fromIndex, toIndex}
	} else if sumExistAt != -1 {
		return []int{sumExistAt, sumExistAt}
	} else {
		return []int{-1}
	}
}
func main() {
	// Example
	sum := 15
	arr := []int{6, 7, 8, 1, 2, 3, 0, 4, 5, 0, 0, 9, 10}
	fmt.Println(findLongestSubarrayBySum(sum, arr)) // [4, 11]
}
