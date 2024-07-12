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
// EXERCISE: Slicing by arguments
//
//   We've a []string, you will get arguments from the command-line,
//   then you will print only the elements that are requested.
//
//   1. Print the []string (it's in the code below)
//
//   2. Get the starting and stopping positions from the command-line
//
//   3. Print the []string slice by slicing it using the starting and stopping
//      positions
//
//   4. Handle the error cases (see the expected output below)
//
//   5. Add new elements to the []string slice literal.
//      Your program should work without changing the rest of the code.
//
//   6. Now, play with your program, get a deeper sense of how the slicing
//      works.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: we expect only two arguments)
//
//  go run main.go 1 2 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: starting index >= 0 && stopping pos <= len(slice) )
//
//  go run main.go -1 5
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  (error because: starting <= stopping)
//
//  go run main.go 3 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  go run main.go 0
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus Warsaw]
//
//
//  go run main.go 1
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus Warsaw]
//
//
//  go run main.go 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Nexus Warsaw]
//
//
//  go run main.go 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Warsaw]
//
//
//  go run main.go 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    []
//
//
//  go run main.go 0 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus]
//
//
//  go run main.go 1 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus]
//
//
//  go run main.go 1 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan]
//
// ---------------------------------------------------------

func main() {
	// uncomment the slice below
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	// Print the slice
	fmt.Printf("%q\n", ships)
	fmt.Println()

	// // If an end isn't passed in, default to the length of the slice.
	// end := len(ships)

	// // ==== Error handling section ====
	// args := os.Args[1:]

	// // No. of arguments must be only 1 or 2
	// if len(args) != 1 && len(args) != 2 {
	// 	fmt.Println("Provide only [starting] and [stopping] positions")
	// 	return
	// }

	// // Get start and end.
	// // If only 1 argument, it is just the start.
	// start, err := strconv.Atoi(args[0])
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }

	// // If 2 arguments, we have the start and end.
	// if len(args) == 2 {
	// 	end, err = strconv.Atoi(args[1])
	// 	if err != nil {
	// 		fmt.Printf("Error: %s", err)
	// 		return
	// 	}
	// }

	// // Start must be >= 0
	// if start < 0 {
	// 	fmt.Println("Wrong positions")
	// 	return
	// }

	// // End must be <= len(ships)
	// if end > len(ships) {
	// 	fmt.Println("Wrong positions")
	// 	return
	// }

	// // Start cannot be greater than end.
	// if start > end {
	// 	fmt.Println("Wrong positions")
	// 	return
	// }
	// // ==== end error handling ====

	// // Print the slice with start and stop
	// fmt.Printf("%s\n", ships[start:end])

	start, end := 0, len(ships)
	var err error

	switch args := os.Args[1:]; len(args) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	// 2 args given: we have a start and end in this case, so grab the end, then fallthrough
	// to the case where we would have only 1 arg and get that as the start.
	case 2:
		end, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Please provide numbers for positions")
			return
		}
		fallthrough
	case 1:
		start, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide numbers for positions")
			return
		}
	}

	// Error case handling
	if l := len(ships); start < 0 || start > l || end > l || start > end {
		fmt.Println("Wrong positions")
		return
	}

	// Print the slice with start and stop
	fmt.Println(ships[start:end])
}
