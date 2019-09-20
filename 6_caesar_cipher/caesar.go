package main

import (
	"fmt"
	"strings"
)

// Caesar Cipher in Golang

var alphabet = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
var shiftKey = 7

// Simple Find index of item in an array function :D
// Wrote it cuz Golang only have the essentials
func findIndex(str string, array *[]string) int {
	var arraytmp = *array
	for index, element := range arraytmp {
		if str == element {
			return index
		}
	}
	return -1
}

func operation(key int, input string) string {
	var allCapsInput = strings.Split(strings.ToUpper(input), "")
	var output string

	for _, char := range allCapsInput {
		var index = findIndex(char, &alphabet)

		if index != -1 {
			output += alphabet[(26+index+key)%26]
		} else {
			output += char
		}
	}
	return output
}

var encrypt = func(text string) string {
	return operation(shiftKey, text)
}

var decrypt = func(text string) string {
	return operation((shiftKey * -1), text)
}

func main() {

	var plaintext = "New York!"
	var ciphertext = encrypt(plaintext)     // "ULD FVYR!"
	var decryptedtext = decrypt(ciphertext) // "NEW YORK!"
	fmt.Println(ciphertext, decryptedtext)
}
