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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const maxTurns = 5

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [guess]")
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Println("Invalid guess - please enter a positive integer")
		return
	}

	// Note that we could just use the global rand.Intn(), which is seeded with a different
	// value on each run, but this way demonstrates how to create our own pseudo-random
	// generator based on Unix timestamps.
	mySrc := rand.NewSource(time.Now().UnixMicro())
	myRnd := rand.New(mySrc)

	for turn := 1; turn <= maxTurns; turn++ {
		n := myRnd.Intn(guess + 1)

		// Uncomment to see random numbers generated
		// fmt.Printf("%d ", n)

		if n != guess {
			continue
		}

		switch myRnd.Intn(3) {
		case 0:
			fmt.Println("You win!")
		case 1:
			fmt.Println("Amazing, you win!")
		case 2:
			fmt.Println("You're a winner!")
		}
		return
	}
	switch myRnd.Intn(3) {
	case 0:
		fmt.Println("You lose...")
	case 1:
		fmt.Println("Too bad, try again")
	case 2:
		fmt.Println("Better luck next time")
	}
}
