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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

const (
	usage = "Usage: [op] [size]"

	validOps        = "*/+-%"
	validOpsMessage = `Valid ops:
  * (Multiply)
  / (Divide)
  + (Add)
  - (Subtract)
  % (Modulo/remainder)`
)

func main() {
	// Size not given
	if len(os.Args) == 2 {
		fmt.Println("Size not given.")
		fmt.Println()
		fmt.Println(usage)
		fmt.Println()
		fmt.Println(validOpsMessage)
		return
	}

	// Invalid no. args given
	if len(os.Args) != 3 {
		fmt.Println(usage)
		fmt.Println()
		fmt.Println(validOpsMessage)
		return
	}

	// Invalid op given
	if !strings.ContainsAny(validOps, os.Args[1]) {
		fmt.Println("Invalid operator.")
		fmt.Println()
		fmt.Println(validOpsMessage)
		return
	}

	op := os.Args[1]

	n, err := strconv.Atoi(os.Args[2])
	if err != nil || n < 1 {
		fmt.Println("Invalid size - must be an integer greater than 0")
		return
	}

	// Print out the header of the table
	fmt.Printf("%5s", op)
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// Rest of the table
	for i := 0; i <= n; i++ {
		// Print out current operand (i.e. the first number being operated on)
		fmt.Printf("%5d", i)
		// Use an inner loop to loop over 0 - n, doing the given operation with each value
		// against the current operand
		var res int
		for j := 0; j <= n; j++ {
			// Do the required operation depending on the passed-in operator
			switch op {
			case "*":
				// Multiply
				res = i * j
			case "+":
				// Add
				res = i + j
			case "-":
				// Subtract
				res = i - j
			case "/", "%":
				// In both "/" (divide) and "%" (modulo/remainder) cases, we need to handle
				// divide-by-zero cases.
				// Any operation trying to divide by zero should just result in 0.
				if j != 0 {
					// Now handle the operator type (divide or modulo/remainder)
					if op == "/" {
						res = i / j
					} else {
						res = i % j
					}
				}
			}
			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}

}
