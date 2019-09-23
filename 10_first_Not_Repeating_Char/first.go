package main

import "fmt"

// First Not Repeating Character - Amazon - Interview Question-
// in Golang
func firstNotRepeatingCharacter(text string) string {
	// Holds repeating and non repeating characters
	var nonRepeatedChars = map[string]int{}
	var RepeatedChars = map[string]string{}

	for index := 0; index < len(text); index++ {
		var charAtIndex = string(text[index])

		var _, isOnRepeatedChars = RepeatedChars[charAtIndex]
		var _, isOnNonRepeatedChars = nonRepeatedChars[charAtIndex]

		if !isOnNonRepeatedChars && !isOnRepeatedChars {
			nonRepeatedChars[charAtIndex] = index
		} else if isOnNonRepeatedChars && !isOnRepeatedChars {
			delete(nonRepeatedChars, charAtIndex)
			RepeatedChars[charAtIndex] = charAtIndex
		}
	}

	// As Map Keys are not ordered in Golang
	// This step is required atm
	var firstChar interface{}
	var tempIndex interface{}
	for key, curIndex := range nonRepeatedChars {
		if firstChar == nil && tempIndex == nil {
			firstChar = key
			tempIndex = curIndex
		} else {
			if tempIndex.(int) > curIndex {
				tempIndex = curIndex
				firstChar = key
			}
		}
	}

	if firstChar == nil {
		return "_"
	}
	return firstChar.(string)
}

func main() {
	input := "thebestshoweverisgameofthrones"
	char := firstNotRepeatingCharacter(input)
	fmt.Println("input: ", input)
	fmt.Println("result: ", char)
}
