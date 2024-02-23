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
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a year number.")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	// Note that the default value of a bool is "false"
	var isLeapYear bool
	if year%400 == 0 {
		isLeapYear = true
	} else if year%100 == 0 {
		isLeapYear = false
	} else if year%4 == 0 {
		isLeapYear = true
	}

	if isLeapYear {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}

	// if year%400 == 0 {
	// 	fmt.Printf("%d is a leap year.\n", year)
	// } else if year%100 == 0 {
	// 	fmt.Printf("%d is not a leap year.\n", year)
	// } else if year%4 == 0 {
	// 	fmt.Printf("%d is a leap year.\n", year)
	// } else {
	// 	fmt.Printf("%d is not a leap year.\n", year)
	// }
}
