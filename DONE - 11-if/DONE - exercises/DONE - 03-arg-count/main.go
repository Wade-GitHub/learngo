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
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	nArgs := len(os.Args) - 1

	if nArgs > 2 {
		fmt.Printf("There are %d arguments\n", nArgs)
	} else if nArgs > 1 {
		fmt.Printf("There are two: \"%s %s\"\n", os.Args[1], os.Args[2])
	} else if nArgs > 0 {
		fmt.Printf("There is one: \"%s\"\n", os.Args[1])
	} else {
		fmt.Println("Give me args")
	}
}
