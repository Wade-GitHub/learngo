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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	usage := "main.go [min] [max]"
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	start, errStart := strconv.Atoi(os.Args[1])
	end, errEnd := strconv.Atoi(os.Args[2])
	if errStart != nil || errEnd != nil {
		fmt.Println("Min and max values must be ints")
		return
	}

	if start > end {
		fmt.Println("Min cannot be greater than max")
		return
	}

	// var total int
	// for i := start; i <= end; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}

	// 	total += i
	// 	fmt.Print(i)
	// 	// If the last number is an odd number, an even i will still be less than it and mess
	// 	// with the "+" printing.
	// 	// Add 1 to i, to make sure the next number (if it's odd) is the end number.
	// 	if i+1 < end {
	// 		fmt.Print(" + ")
	// 	}
	// }

	// for i := start; i <= end; i++ {
	// var total int
	// i := start
	var (
		total int
		i     = start
	)

	for {
		if i > end {
			break
		}

		if i%2 != 0 {
			i++
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
		i++
	}
	fmt.Printf(" = %d\n", total)
}
