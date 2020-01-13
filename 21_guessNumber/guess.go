package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Guess Random Number Game
// In Go

// generated Number won't exceed this value
const maxNum = 500

func main() {

	// Random Number Generator Function
	genRandom := func(from int64, to int64) int64 {
		num, err := rand.Int(rand.Reader, big.NewInt(to-from+1))
		if err != nil {
			fmt.Errorf("An Error Occured %v", err)
			os.Exit(1)
		}
		return num.Int64() + from
	}

	// Reader, read/get user input
	reader := bufio.NewReader(os.Stdin)

	// Range, From & To, a Random Range Between 1, max_num
	rangeFrom := genRandom(0, maxNum)
	rangeTo := genRandom(rangeFrom, maxNum)

	// Random Number that the User have to guess :D
	randomNum := genRandom(rangeFrom, rangeTo)

	// Number of allowed guesses` >= 3
	numberOfGuesses := 3 + (rangeTo-rangeFrom)%3

	var userAttempts int64 // Number of tries by User
	userWon := false       // Is True if User enter random number
	tip := ""              // Tip msg, changes according

	fmt.Printf("Welcome to Guess the Random Number Game :D\n"+
		"Written in Go o_o\n"+
		"The number is between %d ~ %d\n", rangeFrom, rangeTo)
	for numberOfGuesses > userAttempts {
		msg := "\r You have %d tries left,%v Enter a number: "
		fmt.Printf(msg, (numberOfGuesses - userAttempts - 1), tip)

		// User Input
		input, _ := reader.ReadString('\n')
		// Parse Input
		userNumber, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
		userAttempts++

		if userNumber == randomNum {
			userWon = true
			break
		} else if userNumber > randomNum {
			tip = "maybe you should try a smaller number :D"
		} else if userNumber < randomNum {
			tip = "maybe you should try a bigger number :D"
		}
	}

	if userWon {
		fmt.Printf("\nCongrats!!! YOU WON! with %d attempts\n", userAttempts)
	} else {
		fmt.Printf("\n:( You Lost!, you literally had %d attempts, Number was %d\n",
		userAttempts, randomNum)
	}
}
