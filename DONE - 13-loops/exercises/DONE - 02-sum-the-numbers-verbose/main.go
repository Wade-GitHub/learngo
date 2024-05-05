// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	// sum := 0

	// for i := 1; i <= 10; i++ {
	// 	if i == 10 {
	// 		fmt.Print("10 = ")
	// 	} else {
	// 		fmt.Printf("%d + ", i)
	// 	}
	// 	sum += i
	// }
	// fmt.Printf("%d", sum)
	// fmt.Println()

	const start, end = 1, 10

	var sum int
	for i := start; i <= end; i++ {
		sum += i
		fmt.Print(i)
		if i < end {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}
