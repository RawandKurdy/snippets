package main

import "fmt"

// Loops in Golang

func main() {
	var arr = []string{"New York", "Washington", "California"}

	// Foreach loop
	for index, city := range arr {
		fmt.Println(index, city)
	}

	// Traditional for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// While loop // (continuous till condition is true)
	var condition = true
	for condition {
		var city = pop(&arr)
		fmt.Println(city)
		condition = !(len(arr) == 0)
	}
}

// simple pop function
// as golang does only have the important stuff built-in
func pop(arr *[]string) string {
	var arrTmp = *arr
	var indexOfLastItem = len(arrTmp) - 1
	var item = arrTmp[indexOfLastItem]
	*arr = arrTmp[:indexOfLastItem]
	return item
}

// used to show how the while loop would work
