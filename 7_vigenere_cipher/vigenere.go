package main

import (
	"fmt"
	"strings"
)

// Vigenere Cipher in Golang

var alphabet = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
var stringSecret = "Key"                // Secret key used
var secretKeyIndexes []int              // Indexes of letters in string secret
var secretKeyLength = len(stringSecret) // length of the string secret

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

func operation(secretKey string, input string, decrypt bool) string {
	var allCapsInput = strings.Split(strings.ToUpper(input), "")
	var tempCounter int // Temporary Counter
	var output string

	if len(secretKeyIndexes) == 0 { // Convert Secret String into Indexes
		for _, char := range strings.Split(strings.ToUpper(secretKey), "") {
			secretKeyIndexes = append(secretKeyIndexes, findIndex(char, &alphabet))
		}
	}

	for _, char := range allCapsInput {
		var index = findIndex(char, &alphabet)

		if index != -1 {
			var key = secretKeyIndexes[tempCounter%secretKeyLength]
			if decrypt {
				key = key * -1
			}

			tempCounter++
			output += alphabet[(26+index+key)%26]
		} else {
			output += char
		}
	}
	return output
}

var encrypt = func(text string) string {
	return operation(stringSecret, text, false)
}

var decrypt = func(text string) string {
	return operation(stringSecret, text, true)
}

func main() {

	var plaintext = "New York!"
	var ciphertext = encrypt(plaintext)     // "XIU ISPU!"
	var decryptedtext = decrypt(ciphertext) // "NEW YORK!"
	fmt.Println(ciphertext, decryptedtext)
}
