package main

import (
	"fmt"
	"strconv"
)

// is Crypt Solution - Interview Question - Asked by Palantir Technologies
// A Cryptarithm is a mathematical puzzle,
// Here we check if the solution is correct for the crypt puzzle

func isCryptSolution(crypt []string, solution [][]string) bool {
	var letterToNumber = map[string]string{} // map letter > number
	var equation = []string{"", "", ""}      // number1, number2, result

	for _, pair := range solution {
		letter := pair[0]
		number := pair[1]
		letterToNumber[letter] = number
	}

	for wordIndex := 0; wordIndex < len(crypt); wordIndex++ {
		word := crypt[wordIndex] // a word in crypt
		zeroExist := false

		for index := 0; index < len(word); index++ {
			// number that correspond with the letter
			number := letterToNumber[string(word[index])]

			if zeroExist { // if zero exist at start then exit
				return false
			}
			if number == "0" && index == 0 {
				zeroExist = true
			}
			equation[wordIndex] += number
		}
	}
	// As usual GoLang makes stuff fun :D
	// I didnt want to repeat this 3 times so why not make
	// A mini function :)
	parseInt := func(str string) int64 {
		number, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return number
		}
		panic(err)
	}

	// Evaluate the equation and see if its valid
	return parseInt(equation[0])+parseInt(equation[1]) == parseInt(equation[2])
}

func main() {
	var crypt = []string{"SEND",
		"MORE",
		"MONEY"}
	var solution = [][]string{{"O", "0"},
		{"M", "1"},
		{"Y", "2"},
		{"E", "5"},
		{"N", "6"},
		{"D", "7"},
		{"R", "8"},
		{"S", "9"}}

	fmt.Println(isCryptSolution(crypt, solution)) // True
}
