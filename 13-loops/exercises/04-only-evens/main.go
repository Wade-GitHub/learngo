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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	usage := "main.go [min] [max]"
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Min value must be an int")
		return
	}

	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Max value must be an int")
		return
	}

	var total int

	for i := start; i <= end; i++ {
		if i%2 != 0 {
			continue
		}

		total += i
		fmt.Print(i)
		// If the last number is an odd number, an even i will still be less than it and mess
		// with the "+" printing.
		// Add 1 to i, to make sure the next number (if it's odd) is the end number.
		if i+1 < end {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", total)
}
