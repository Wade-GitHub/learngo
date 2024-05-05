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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
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

	// Note that we could just use the global rand.Intn(), which is seeded with a different
	// value on each run, but this way demonstrates how to create our own pseudo-random
	// generator based on Unix timestamps.
	mySrc := rand.NewSource(time.Now().UnixMicro())
	myRnd := rand.New(mySrc)

	for turn := 1; turn <= maxTurns; turn++ {
		n := myRnd.Intn(guess + 1)

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
