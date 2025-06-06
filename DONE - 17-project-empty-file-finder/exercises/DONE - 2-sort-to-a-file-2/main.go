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
	"slices"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file with their ordinals
//
//  Use the previous exercise.
//
//  This time, print the sorted items with ordinals
//  (see the expected output)
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     1. apple
//     2. banana
//     3. orange
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + You can use strconv.AppendInt function to append an int
//     to a byte slice. strconv contains a lot of functions for appending
//     other basic types to []byte slices as well.
//
//   + You can append individual characters to a byte slice using
//     rune literals (because: rune literal are typeless numerics):
//
//     var slice []byte
//     slice = append(slice, 'h', 'i', ' ', '!')
//     fmt.Printf("%s\n", slice)
//
//     Above code prints: hi !
// ---------------------------------------------------------

// func main() {
// 	items := os.Args[1:]
// 	if len(items) == 0 {
// 		fmt.Println("Send me some items and I will sort them")
// 		return
// 	}

// 	sort.Strings(items)

// 	var data []byte
// 	for _, s := range items {
// 		data = append(data, s...)
// 		data = append(data, '\n')
// 	}

// 	err := ioutil.WriteFile("sorted.txt", data, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	// Chop off the first arg, the exe name.
	args := os.Args[1:]

	// Sort the args
	slices.Sort(args)

	// Slice of bytes to write to file
	var data []byte
	// strconv.AppendInt()

	// Append each arg to the data slice.
	for i, arg := range args {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, arg...)
		data = append(data, '\n')
	}

	err := os.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	// fmt.Printf("%s\n", data)

}
