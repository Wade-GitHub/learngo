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
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

func main() {
	expected := strings.Fields("yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday")
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// ...

	// Actual code:
	lyric = append([]string{"yesterday"}, lyric...)
	lyric = append(lyric, lyric[8:13]...)
	lyric = append(lyric[0:8], lyric[13:]...)

	fmt.Printf("Expected: %s\n", expected)
	fmt.Printf("Result:   %s\n", lyric)

	// Visualisation of steps:
	// fmt.Print(strings.Repeat(" ", 11))
	// for i := 0; i < len(lyric); i++ {
	// 	fmt.Printf("%-11d", i)
	// }
	// fmt.Println()
	// fmt.Printf("Expected: %-10s\n", expected)
	// fmt.Printf("Result:   %-10s\n", lyric)

	// fmt.Println("====")

	// fmt.Printf("lyric len: %d, lyric cap: %d\n", len(lyric), cap(lyric))
	// lyric = append([]string{"yesterday"}, lyric...)
	// fmt.Printf("lyric len: %d, lyric cap: %d\n", len(lyric), cap(lyric))

	// fmt.Print(strings.Repeat(" ", 11))
	// for i := 0; i < len(lyric); i++ {
	// 	fmt.Printf("%-11d", i)
	// }
	// fmt.Println()
	// fmt.Printf("Expected: %-10s\n", expected)
	// fmt.Printf("Result:   %-10s\n", lyric)

	// fmt.Println("====")

	// fmt.Printf("lyric len: %d, lyric cap: %d\n", len(lyric), cap(lyric))
	// lyric = append(lyric, lyric[8:13]...)
	// fmt.Printf("lyric len: %d, lyric cap: %d\n", len(lyric), cap(lyric))

	// fmt.Print(strings.Repeat(" ", 11))
	// for i := 0; i < len(lyric); i++ {
	// 	fmt.Printf("%-11d", i)
	// }
	// fmt.Println()
	// fmt.Printf("Expected: %-10s\n", expected)
	// fmt.Printf("Result:   %-10s\n", lyric)

	// fmt.Println("====")

	// fmt.Printf("lyric len: %d, lyric cap: %d\n", len(lyric), cap(lyric))
	// lyric = append(lyric[0:8], lyric[13:]...)
	// fmt.Printf("lyric len: %d, lyric cap: %d\n", len(lyric), cap(lyric))

	// fmt.Print(strings.Repeat(" ", 11))
	// for i := 0; i < len(lyric); i++ {
	// 	fmt.Printf("%-11d", i)
	// }
	// fmt.Println()
	// fmt.Printf("Expected: %-10s\n", expected)
	// fmt.Printf("Result:   %-10s\n", lyric)
}
