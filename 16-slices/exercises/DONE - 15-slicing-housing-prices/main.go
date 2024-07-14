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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slicing the Housing Prices
//
//  We have received housing prices. Your task is to print only the requested
//  columns of data (see the expected output).
//
//  1. Separate the data by the newline ("\n") -> rows
//
//  2. Separate each row of the data by the separator (",") -> columns
//
//  3. Find the from and to positions inside the columns depending
//     on the command-line arguments.
//
//  4. Print only the requested column headers and their data
//
//
// RESTRICTIONS
//
//  + You should use slicing when printing the columns.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  go run main.go Location
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  NOTE : Finds the position of the Size column and slices the columns
//         appropriately.
//
//  go run main.go Size
//    Size           Beds           Baths          Price
//
//    100            2              1              100000
//    150            3              2              200000
//    200            4              3              400000
//    500            10             5              1000000
//
//
//  NOTE : Finds the positions of the Size and Baths columns and
//         slices the columns appropriately.
//
//  go run main.go Size Baths
//    Size           Beds           Baths
//
//    100            2              1
//    150            3              2
//    200            4              3
//    500            10             5
//
//
//  go run main.go Beds Price
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  Note : It works even if the given column name does not exist.
//
//  go run main.go Beds NotExist
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  go run main.go NotExist NotExist
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
// Note : It works even if the Price's index > Size's index
//
//        In that case, it resets the invalid starting position to 0.
//
//        But it still uses the Size column.
//
//  go run main.go Price Size
//    Location       Size
//
//    New York       100
//    New York       150
//    Paris          200
//    Istanbul       500
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// Split all string data into a slice so we can get the relevant fields out as slices.
	dataSlice := strings.Split(data, "\n")
	headings := strings.Split(dataSlice[0], separator)

	var rows [][]string
	for _, line := range dataSlice[1:] {
		rows = append(rows, strings.Split(line, separator))
	}

	// Set the initial "from" and "to" indexes, in the case fewer than 2 args are passed.
	from, to := 0, len(headings)

	// Getting args from CLI.
	switch args := os.Args[1:]; len(args) {
	default:
		fallthrough
	case 0:
		// No args passed in - break from the switch statement. Default "from" and "to" will
		// be used.
		break
	case 2:
		// See if arg2 exists in the slice.
		to = slices.Index(headings, args[1])
		if to == -1 {
			// If not, set "to" to the length of the slice.
			to = len(headings)
		} else {
			// If the arg exists in the slice, we need to add 1 to the "to" index, or it will
			// exclude that column in the table when printing it.
			to += 1
		}
		// Fall through to pick up the "from" argument.
		fallthrough
	case 1:
		// See if arg1 exists in the slice.
		from = slices.Index(headings, args[0])
		if from == -1 {
			// If not, set it to the beginning.
			from = 0
		}
	}

	// If the "from" index is greater than the "to" index, set "from" to the start of the
	// slice, since it's invalid.
	if from > to {
		from = 0
	}

	// Print the table headings.
	for _, heading := range headings[from:to] {
		fmt.Printf("%-15s", heading)
	}
	fmt.Println()
	fmt.Println()

	// Print the table data.
	for _, row := range rows {
		for _, item := range row[from:to] {
			fmt.Printf("%-15s", item)
		}
		fmt.Println()
	}
}
