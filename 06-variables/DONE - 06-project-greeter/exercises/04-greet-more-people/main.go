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
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	// fmt.Println("There are", len(os.Args)-1, "people!")
	// fmt.Println("Hello great", os.Args[1], "!")
	// fmt.Println("Hello great", os.Args[2], "!")
	// fmt.Println("Hello great", os.Args[3], "!")
	// fmt.Println("Nice to meet you all.")

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
	nameCount := len(os.Args) - 1

	// printPeopleCount(nameCount)
	// if nameCount == 1 {
	// 	fmt.Println("There is", nameCount, "person!")
	// } else {
	// 	fmt.Println("There are", nameCount, "people!")
	// }
	if nameCount != 1 {
		fmt.Println("There are", nameCount, "people!")
	} else {
		fmt.Println("There is 1 person!")
	}

	// greetPeople(os.Args)
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Hello great", os.Args[i], "!")
	}

	// printSalutations(nameCount)
	if nameCount > 1 {
		fmt.Println("Nice to meet you all.")
	} else if nameCount > 0 {
		fmt.Println("Nice to meet you.")
	}
}
