package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getting an input from the console
// in Golang

func main() {
	reader := bufio.NewReader(os.Stdin) // reads from stdin (input)
	fmt.Print("What is your name?")     // prompt msg

	line, _ := reader.ReadString('\n') // reads string till new line occur

	// I used print format and I removed new line separator from the input
	fmt.Printf("Hello %s :) \n", strings.Replace(line, "\n", "", -1))
}
