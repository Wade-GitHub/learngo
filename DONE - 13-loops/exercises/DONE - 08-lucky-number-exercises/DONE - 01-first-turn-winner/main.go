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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

const maxTurns = 5

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: [guess]")
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Println("Invalid guess - please enter a positive integer")
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

		// if n == guess {
		// 	if turn == 0 {
		// 		fmt.Println("First try! Well done!")
		// 	} else {
		// 		fmt.Println("You win!")
		// 	}
		// 	return
		// }

		if n != guess {
			continue
		}

		if turn == 1 {
			fmt.Println("First try! Well done!")
		} else {
			fmt.Println("You win!")
		}
		return
	}
	fmt.Println("You lose...")
}
