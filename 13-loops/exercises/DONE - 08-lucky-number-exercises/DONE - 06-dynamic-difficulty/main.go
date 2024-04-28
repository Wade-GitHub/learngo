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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

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

	// Difficulty adjustment: if user passes in a guess <10, the minimum guess range and turns
	// for the game is 10.
	// If user passes in >10, use that number for the guess range and adjust no. of turns via
	// a rudimentary formula: (guess / 3) + 8.
	// This is to make sure turns are adjusted to the size of the guess (i.e. larger guesses
	// should allow more turns for fairness)
	minGuessRange, turns := 10, 10
	if guess > minGuessRange {
		minGuessRange = guess
		turns = (guess / 3) + 8
	}

	// Note that we could just use the global rand.Intn(), which is seeded with a different
	// value on each run, but this way demonstrates how to create our own pseudo-random
	// generator based on Unix timestamps.
	mySrc := rand.NewSource(time.Now().UnixMicro())
	myRnd := rand.New(mySrc)

	for turn := 1; turn <= turns; turn++ {
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
