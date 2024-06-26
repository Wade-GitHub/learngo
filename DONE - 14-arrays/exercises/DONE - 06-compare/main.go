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
// EXERCISE: Compare the Arrays
//
//  1. Uncomment the code
//
//  2. Fix the problems so that arrays become comparable
//
// EXPECTED OUTPUT
//  true
//  true
//  false
// ---------------------------------------------------------

func main() {
	week := [2]string{"Monday", "Tuesday"}
	wend := [2]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [5]int{2, 4, 6, 8, 10}
	evens2 := [5]int{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]uint8{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
