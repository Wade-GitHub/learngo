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
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if l := len(args); l == 0 || l > 5 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	var nums [5]int
	for i := range args {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			continue
		}
		nums[i] = n
	}

	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] > nums[j] {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 		}
	// 		// fmt.Printf("%d\n", nums)
	// 	}
	// }

	// Another way of doing bubble sort:
	for range nums {
		for i, n := range nums {
			if i < len(nums)-1 && n > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	fmt.Printf("%d\n", nums)
}
