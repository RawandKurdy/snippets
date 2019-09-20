package main

import "fmt"

// Functions in Golang

// 1- Traditional function
// also used to demonstrate how an anonymous func can be useful
// It runs a function passed as a param
func traditionalFunc(anotherFunc funcType) {
	anotherFunc()
}

// Custom type that uses func as base
// to allow function as params
type funcType func()

func main() {
	// 2- Anonymous Function
	// I passed it as a parameter
	traditionalFunc(func() {
		fmt.Println("Hey World!")
	})

	// 3- Function Expression
	var anotherFunc = func() {
		fmt.Println("Hey World Again!")
	}

	// Exec function expr.
	anotherFunc()
}
