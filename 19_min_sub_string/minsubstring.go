package main

import "fmt"

// Minimum Substring With All Chars - Interview Question -
// Asked by Facebook
// In Golang
func minSubstringWithAllChars(s string, t string) string {
	var size = len(s) + 1 // length of the text
	for x := 0; x < size; x++ {
		for p := 0; p < size-x; p++ {
			c := "" // sub string candidate
			// also gets filtered
			// chars that are not in t get removed for now
			for _, letter := range s[p : p+x] {
				if containsRune(t, letter) {
					c += string(letter)
				}
			}

			// if sub string contain all characters that are in t
			// then correct answer is found :D
			if len(setEmu(c)) == len(t) {
				return s[p : p+x]
			}
		}
	}
	return ""
}

// As Golang does not support sets
// We use a Map to have a duplicate free collection
func setEmu(s string) map[string]bool {
	m := map[string]bool{}
	for _, letter := range s {
		m[string(letter)] = true
	}
	return m
}

// Golang supports has this by default
// but Codesignal didn't provide a way
// to import the strings package
// used to help in filtering the string
func containsRune(t string, l rune) bool {
	for _, letter := range t {
		if letter == l {
			return true
		}
	}
	return false
}

func main() {
	// Example
	s := "adobecodebanc"
	t := "abc"
	fmt.Println(minSubstringWithAllChars(s, t)) // "banc"
}
