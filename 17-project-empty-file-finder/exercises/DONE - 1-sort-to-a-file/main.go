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
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
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
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

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

	// Append each arg to the data slice.
	for _, arg := range args {
		data = append(data, arg...)
		data = append(data, '\n')
	}

	err := os.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

}
