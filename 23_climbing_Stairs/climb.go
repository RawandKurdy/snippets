package main

import "fmt"

// Climbing Stairs - Interview Question -
// Asked By Apple and Adobe
// in Go

func climbingStairs(n int) int {
	r := []int{1, 1, 2} // Initial Result Set

	if n <= 2 {
		return r[n]
	}

	for step := 3; step < n+1; step++ {
		r = append(r, r[step-1]+r[step-2])
	}
	return r[n]
}

func main() {
	n := 26 // Steps
	ways := climbingStairs(n)
	fmt.Println(ways) // 196418
}
