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
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

const maxTurns = 5

func main() {
	args := os.Args[1:]
	nArgs := len(args)

	// -v flag is optional, so acceptable no. of args is either 1 (if just passing the guess)
	// or 2 (if passing a guess + the -v flag)
	if nArgs < 1 || nArgs > 2 {
		fmt.Println("Usage: [guess] [-v]")
		return
	}

	// Declare some variables we'll need to work with ahead of time
	var (
		verboseMode bool

		guess int
		err   error
	)

	// Go through args.
	// If "-v" arg exists, set verboseMode.
	// Anything else *should* be the user's guess. Handle it as usual.
	for _, arg := range args {
		switch arg {
		case "-v":
			verboseMode = true
		default:
			guess, err = strconv.Atoi(arg)
			if err != nil || guess < 0 {
				fmt.Println("Invalid guess - please enter a positive integer for guess")
				return
			}
		}
	}

	// If user passes in a number less than 10, make the minimum guess range 10.
	minGuessRange := 10
	if guess > minGuessRange {
		minGuessRange = guess
	}

	// Note that we could just use the global rand.Intn(), which is seeded with a different
	// value on each run, but this way demonstrates how to create our own pseudo-random
	// generator based on Unix timestamps.
	mySrc := rand.NewSource(time.Now().UnixMicro())
	myRnd := rand.New(mySrc)

	for turn := 1; turn <= maxTurns; turn++ {
		// n := myRnd.Intn(guess + 1)
		n := myRnd.Intn(minGuessRange + 1)

		// Print each number being compared to guess if verboseMode is on
		if verboseMode {
			fmt.Printf("%d ", n)
		}

		if n != guess {
			continue
		}

		fmt.Println("You win!")
		return
	}
	fmt.Println("You lose...")
}
