package main

import "fmt"

// Amend The Sentence - Adobe - Interview Question -
// in Golang

func amendTheSentence(text string) string {
	var newStr string

	for index, char := range text {
		if 'A' <= char && char <= 'Z' {
			if index != 0 {
				newStr += " "
			}
			newStr += string(char + 32)
		} else {
			newStr += string(char)
		}
	}
	return newStr
}

func main() {
	input := "YouAreInTheRightPlace"
	fmt.Println("input: ", input)
	fmt.Println("output: ", amendTheSentence(input))
}
