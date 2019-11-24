package main

import "fmt"

// DivNumber Challenge
// Given three integers k, l and r
// return the number of integers between l and r inclusive
// that have exactly k divisors.
// In Golang

func main() {
	// Example
	k := 3
	l := 2
	r := 49
	fmt.Println(divNumber(k, l, r)) // 4

	// We have 4 integers in that range
	// that have exactly 3 divisors - 4, 9, 25, 49.
}

func divNumber(k int, l int, r int) int {
	// di is the number of Numbers
	// that has 'k' exact divisors
	di := 0
	for l <= r {
		//d is number of divisors of individual numbers
		d := 0
		// c is counter(Number)
		c := 1

		for c <= l {
			if l%c == 0 {
				d++
			}
			c++
		}
		if d == k {
			di++
		}
		l++
	}
	return di
}
