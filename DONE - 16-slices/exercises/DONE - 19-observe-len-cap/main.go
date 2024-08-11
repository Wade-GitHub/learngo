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
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	//
	// 2. print the length and capacity of the games slice
	//
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//
	// 4. print the length and capacity of the games slice
	//
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//
	// 6. print the length and capacity of the games slice
	//
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)

	// var games []string
	// fmt.Printf("games length: %d, games capacity: %d", len(games), cap(games))
	// games := []string{}
	// fmt.Printf("games length: %d, games capacity: %d\n", len(games), cap(games))

	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("games length: %d, games capacity: %d\n", len(games), cap(games))
	games := []string{"pacman", "mario", "tetris", "doom"}

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	// for ... {
	// 	fmt.Printf("games[:%d]'s len: %d cap: %d\n", ...)
	// }
	for i := 0; i <= 4; i++ {
		s := games[:i]
		fmt.Printf("s length: %d, s capacity: %d\n", len(s), cap(s))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	//
	// 2. print the games and the new slice's len and cap
	//
	// 3. append a new element to the new slice
	//
	// 4. print the new slice's lens and caps
	//
	// 5. repeat the last two steps 5 times (use a loop)
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	// []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	fmt.Println()

	// zero := ...
	// fmt.Printf("games's     len: %d cap: %d\n", ...)
	// fmt.Printf("zero's      len: %d cap: %d\n", ...)

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }
	zero := games[:0]
	fmt.Println("zero address:", &zero)
	fmt.Println("games address:", &games)
	fmt.Println("zero original:", len(zero))
	fmt.Printf("games length: %d, games capacity: %d\n", len(games), cap(games))
	fmt.Printf("zero length: %d, zero capacity: %d\n", len(zero), cap(zero))

	zero = append(zero, "zelda")
	fmt.Printf("zero length: %d, zero capacity: %d\n", len(zero), cap(zero))
	fmt.Println("---")
	fmt.Printf("games: %q\n", games)
	fmt.Printf("zero: %q\n", zero)
	fmt.Println("---")

	for _, item := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, item)
		fmt.Printf("zero length: %d, zero capacity: %d\n", len(zero), cap(zero))
	}
	fmt.Println("---")
	fmt.Printf("games: %q\n", games)
	fmt.Printf("zero: %q\n", zero)
	fmt.Println("---")
	fmt.Printf("games length: %d, games capacity: %d\n", len(games), cap(games))

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	// for ... {
	//   s := zero[:n]
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", ...)
	// }
	for n := range zero {
		s := zero[:n]
		fmt.Printf("s length: %d, s capacity: %d\n", len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	// zero = ...
	// for ... {
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", ...)
	// }
	fmt.Println("zero before:", len(zero))
	zero = zero[:cap(zero)]
	fmt.Println("zero after:", len(zero))
	for n := range zero {
		s := zero[:n]
		fmt.Printf("s length: %d, s capacity: %d\n", len(s), cap(s))
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	//
	// 2. change the same element of the games slice
	//
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	// fmt.Printf("zero  : %q\n", zero)
	// fmt.Printf("games : %q\n", games)
	zero[2] = "foo"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games  : %q\n", games)

	games[2] = "bar"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games  : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	s := games[:cap(games)+1]
	fmt.Println(s)
}
