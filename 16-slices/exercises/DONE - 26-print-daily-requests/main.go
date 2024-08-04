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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
//      reqs := []int{
// 	      500, 600, 250,
// 	      200, 400, 50,
// 	      900, 800, 600,
// 	      750, 250, 100,
// 	      150, 654, 235,
// 	      320, 534, 765,
// 	      121, 876, 285,
// 	      543, 642,
// 	      // the last element is missing (your code should be able to handle this)
// 	      // that is why you shouldn't calculate the `size` below manually.
//      }
//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	// reqs := []int{
	// 	500, 600, 250, // 1st day: 1350 requests
	// 	200, 400, 50, // 2nd day: 650 requests
	// 	900, 800, 600, // 3rd day: 2300 requests
	// 	750, 250, 100, // 4th day: 1100 requests
	// 	// grand total: 5400 requests
	// }

	reqs := []int{
		500, 600, 250,
		200, 400, 50,
		900, 800, 600,
		750, 250, 100,
		150, 654, 235,
		320, 534, 765,
		121, 876, 285,
		543, 642,
		// the last element is missing (your code should be able to handle this)
		// that is why you shouldn't calculate the `size` below manually.
	}

	// ================================================
	// #1: Make a new slice with the exact size needed.

	// _ = reqs // remove this when you start

	// size := 0 // you need to find the size.
	// Each 24-hour day is made up of 3 8-hour data points, so we must divide the length of the
	// logs into 3 (or "N") to get the 8-hour blocks.
	l := len(reqs)
	size := l / N
	if l%N != 0 {
		// If the length of the requests isn't divisible by 3 (N), increase the size to account
		// for an incomplete 8-hour block.
		size++
	}
	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	// _ = daily // remove this when you start
	// var day []int
	// for i := 0; i < len(reqs); i += N {
	// 	if i+N > len(reqs) {
	// 		day = reqs[i:]
	// 	} else {
	// 		day = reqs[i : i+N]
	// 	}
	// 	daily = append(daily, day)
	// }
	for N <= len(reqs) {
		daily = append(daily, reqs[:N])
		reqs = reqs[N:]
	}

	if len(reqs) > 0 {
		daily = append(daily, reqs)
	}

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	// ...
	var grandTotal int
	for i, day := range daily {
		var total int
		for _, req := range day {
			fmt.Printf("%-10d %d\n", i+1, req)
			total += req
		}
		fmt.Printf("%9s: %d\n", "TOTAL", total)
		fmt.Println()
		grandTotal += total
	}
	fmt.Printf("%9s: %d\n", "GRAND", grandTotal)

	// ================================================
}
