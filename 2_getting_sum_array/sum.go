// Getting the sum of numbers of an array
// in Golang
package main

import "fmt"

func main() {
	var arr = []int{1, 22, 3, 44, 5, 66, 7}

	// Using for loop
	var sum int
	for _, val := range arr {
		sum += val
	}

	fmt.Println(sum)
}
