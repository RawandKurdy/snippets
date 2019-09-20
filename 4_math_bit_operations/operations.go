package main

import (
	"fmt"
	"math"
)

// Math and Bit Operation
// in Golang
func main() {
	// Multiplication
	var resultMul = 3 * 3
	fmt.Println(resultMul) // 9

	// Power
	var resultPow = math.Pow(3, 3)
	fmt.Println(resultPow) // 27

	// Addition
	var resultAdd = 3 + 3
	fmt.Println(resultAdd) // 6

	// Subtraction
	var resultSub = 3 - 3
	fmt.Println(resultSub) // 0

	// Division
	var resultDiv = 3 / 3
	fmt.Println(resultDiv) // 1

	// Bitwise Right Shift
	var resultBitR = 3 >> 3
	fmt.Println(resultBitR) // 0

	// Bitwise Left Shift
	var resultBitL = 3 << 3
	fmt.Println(resultBitL) // 24

	// Bitwise OR
	var resultBitOr = 3 | 9
	fmt.Println(resultBitOr) // 11

	// Bitwise AND
	var resultBitAnd = 3 & 9
	fmt.Println(resultBitAnd) // 1
}
