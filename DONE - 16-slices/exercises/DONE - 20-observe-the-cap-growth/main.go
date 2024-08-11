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
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 1e7 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

func main() {
	// var (
	// 	s                 []int
	// 	l, newCap, oldCap int
	// 	ratio             float64
	// )

	// // Get all the values for the nil slice and print.
	// // Note that "oldCap" and "newCap" will be the same here.
	// l = len(s)
	// oldCap = cap(s)
	// newCap = cap(s)
	// ratio = float64(newCap) / float64(oldCap)
	// fmt.Printf("len:%-15d cap:%-15d growth:%.2f\n", l, newCap, ratio)

	// // Loop through to 1e7.
	// for i := 0; i < 1e7; i++ {
	// 	// Get the old capacity, append an item to the slice, then get the new length and
	// 	// capacities, and the new ratio, to print.
	// 	oldCap = cap(s)

	// 	s = append(s, i)
	// 	l = len(s)
	// 	newCap = cap(s)

	// 	if newCap == oldCap {
	// 		// If the capacity of the slice didn't change, continue on with the loop.
	// 		continue
	// 	}
	// 	ratio = float64(newCap) / float64(oldCap)
	// 	fmt.Printf("len:%-15d cap:%-15d growth:%.2f\n", l, newCap, ratio)
	// }

	var (
		nums   []int
		oldCap float64
	)

	for len(nums) < 1e7 {
		c := float64(cap(nums))

		if c == 0 || c != oldCap {
			fmt.Printf("len:%-15d cap:%-15g growth:%.2f\n", len(nums), c, c/oldCap)
		}

		oldCap = c
		nums = append(nums, 1)
	}
}
