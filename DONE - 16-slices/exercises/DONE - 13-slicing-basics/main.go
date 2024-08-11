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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	var (
		nums,
		evens,
		odds,
		middle,
		first2,
		last2,
		evensLast1,
		oddsLast2 []int
	)

	strNums := strings.Split(data, " ")
	for i := range strNums {
		n, err := strconv.Atoi(strNums[i])
		if err != nil {
			fmt.Println("Invalid data")
		}
		nums = append(nums, n)
	}

	// Collect the even numbers
	// evens = append(evens, nums[0], nums[1], nums[2])
	evens = append(evens, nums[:3]...)

	// Collect the odd numbers
	// odds = append(odds, nums[3], nums[4], nums[5])
	odds = append(odds, nums[3:]...)

	// Collect the two numbers in the middle
	// middle = append(middle, nums[2], nums[3])
	middle = append(middle, nums[2:3]...)

	// Collect the first 2 numbers
	first2 = append(first2, nums[0:2]...)

	// Collect the last 2 numbers
	last2 = append(last2, nums[len(nums)-2:]...)

	// Slice the evens for the last number
	evensLast1 = append(evensLast1, evens[len(evens)-1:]...)

	// Slice the odds for the last two numbers
	oddsLast2 = append(oddsLast2, odds[len(odds)-2:]...)

	fmt.Printf("nums         : %d\n", nums)
	fmt.Printf("evens        : %d\n", evens)
	fmt.Printf("odds         : %d\n", odds)
	fmt.Printf("middle       : %d\n", middle)
	fmt.Printf("first 2      : %d\n", first2)
	fmt.Printf("last 2       : %d\n", last2)
	fmt.Printf("evens last 1 : %d\n", evensLast1)
	fmt.Printf("odds last 2  : %d\n", oddsLast2)
}
