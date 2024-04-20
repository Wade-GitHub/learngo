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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------

func main() {
	// Print usage if user doesn't pass anything to program.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [size]")
		return
	}

	s, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: invalid size. Must be an integer.")
		return
	}

	if s < 0 {
		fmt.Println("Error: size cannot be negative.")
		return
	}

	// Print the operator
	fmt.Printf("%5s", "X")

	// Print the heading.
	// The heading trails after the operator, and is every number up to "s".
	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// For each number up to "s", print the number at the start.
	// This sets up the column of numbers up to "s".
	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)

		// This inner loop works similar to the first loop we did earlier.
		// Go through every number up to "s", and multiply it by the iteration we're on in the
		// outer loop.
		for j := 0; j <= s; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
