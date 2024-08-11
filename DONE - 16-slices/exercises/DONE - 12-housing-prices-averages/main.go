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
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locations                            []string
		sizes, beds, baths, prices           []int
		avgSize, avgBeds, avgBaths, avgPrice float64
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		rowData := strings.Split(row, separator)

		// Extract location.
		loc := rowData[0]

		// Extract size.
		size, err := strconv.Atoi(rowData[1])
		if err != nil {
			fmt.Println("Invalid size")
			return
		}

		// Extract bed count.
		bed, err := strconv.Atoi(rowData[2])
		if err != nil {
			fmt.Println("Invalid bed count")
			return
		}

		// Extract bath count.
		bath, err := strconv.Atoi(rowData[3])
		if err != nil {
			fmt.Println("Invalid bath count")
			return
		}

		// Extract price.
		price, err := strconv.Atoi(rowData[4])
		if err != nil {
			fmt.Println("Invalid price")
			return
		}

		// Append data into slices.
		locations = append(locations, loc)
		sizes = append(sizes, size)
		beds = append(beds, bed)
		baths = append(baths, bath)
		prices = append(prices, price)
	}

	// Calculate the averages for the following for each row:
	// - Size
	// - Beds
	// - Baths
	// - Price
	for i := 0; i < len(rows); i++ {
		avgSize += float64(sizes[i])
		avgBeds += float64(beds[i])
		avgBaths += float64(baths[i])
		avgPrice += float64(prices[i])
	}
	avgSize /= float64(len(sizes))
	avgBeds /= float64(len(beds))
	avgBaths /= float64(len(baths))
	avgPrice /= float64(len(prices))

	// Print the header.
	for _, item := range strings.Split(header, separator) {
		fmt.Printf("%-15s", item)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 75))

	// Print each row of data.
	for i := 0; i < len(rows); i++ {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}
	fmt.Println(strings.Repeat("=", 75))

	// Print the averages.
	fmt.Print(strings.Repeat(" ", 15))
	fmt.Printf("%-15.2f", avgSize)
	fmt.Printf("%-15.2f", avgBeds)
	fmt.Printf("%-15.2f", avgBaths)
	fmt.Printf("%-15.2f", avgPrice)
	fmt.Println()
}
