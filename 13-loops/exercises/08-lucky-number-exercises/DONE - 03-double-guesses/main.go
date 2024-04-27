// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurns     = 5
	invalidGuess = "Invalid guess - please enter a positive integer for both guesses"
)

func main() {
	nArgs := len(os.Args[1:])

	// Should only be 1 or 2 guesses
	if nArgs < 1 || nArgs > 2 {
		fmt.Println("Usage: go run main.go [guess1] [guess2]")
		return
	}

	// Get the first guess
	guess1, err := strconv.Atoi(os.Args[1])
	if err != nil || guess1 <= 0 {
		fmt.Println(invalidGuess)
		return
	}

	// If 2 guesses passed in, also get the second guess
	var guess2 int
	if nArgs == 2 {
		guess2, err = strconv.Atoi(os.Args[2])
		if err != nil || guess2 <= 0 {
			fmt.Println(invalidGuess)
			return
		}
	}

	// If guess2 is greater than guess1, our max guesses are guess2.
	// Otherwise, max guesses are guess1.
	maxGuess := guess1
	if guess2 > guess1 {
		maxGuess = guess2
	}

	// Note that we could just use the global rand.Intn(), which is seeded with a different
	// value on each run, but this way demonstrates how to create our own pseudo-random
	// generator based on Unix timestamps.
	mySrc := rand.NewSource(time.Now().UnixMicro())
	myRnd := rand.New(mySrc)

	for turn := 1; turn <= maxTurns; turn++ {
		// Since guess2 defaults to 0 if no second guess was passed in, we have to add 1 to
		// the int returned by Intn, since it generates a number between 0 and maxGuess.
		// This means the user's guess cannot be 0.
		n := myRnd.Intn(maxGuess) + 1

		// Uncomment to see random numbers generated
		// fmt.Printf("%d ", n)

		if n != guess1 && n != guess2 {
			continue
		}

		fmt.Println("You win!")
		return
	}
	fmt.Println("You lose...")
}
