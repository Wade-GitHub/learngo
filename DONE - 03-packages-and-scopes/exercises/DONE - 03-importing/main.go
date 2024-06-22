// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

// ?
// ?
// ?
import (
	fmt1 "fmt"
	fmt2 "fmt"
	fmt3 "fmt"
	// fmt2 "fmt"
	// fmt3 "fmt"
)

func main() {
	// ?
	fmt1.Println("Hello")
	// ?
	fmt2.Println("Hey")
	// ?
	fmt3.Println("Hi")
}