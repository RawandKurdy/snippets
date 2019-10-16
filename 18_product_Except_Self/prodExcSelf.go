package main

import "fmt"

// Product Except Self [in O(n)] - Interview Question -
// Asked by Amazon, LinkedIn, Facebook, Microsoft & Apple
// In Golang

func productExceptSelf(nums []int, m int) int {
	s := 0 // Sum
	p := 1 // Product
	for _, num := range nums {
		s = (p + s*num) % m
		p = (p * num) % m
	}
	return s
}

func main() {
	// Example
	nums := []int{3, 3, 9, 5, 5, 4, 2, 8, 5, 9} // Input Array
	m := 17                                     // Mod value

	fmt.Println(productExceptSelf(nums, m)) // 16
}
