package main

import "fmt"

// This is how you create an array (Slice)
// in Golang
func main() {
	var arr = []int{1, 2, 3}
	fmt.Println(arr)

	// to add a value
	arr = append(arr, 4) // to the end
	arr = append([]int{0}, arr...)
	fmt.Println(arr)

	// to remove a value
	arr = arr[:len(arr)-1] // last element
	arr = arr[1:]          // first element
	fmt.Println(arr)
}
