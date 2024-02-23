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
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Give me a month name.")
	// 	return
	// }

	// var isLeapYear bool

	// year := time.Now().Year()
	// It is a leap year if the year is:
	// - divisible by 400, or
	// - divisible by 4, but not 100
	// if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
	// 	isLeapYear = true
	// } else {
	// 	isLeapYear = false
	// }

	// month := strings.ToLower(os.Args[1])

	// if month == "january" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else if month == "february" {
	// 	if isLeapYear {
	// 		fmt.Printf("%q has 29 days.\n", month)
	// 	} else {
	// 		fmt.Printf("%q has 28 days.\n", month)
	// 	}
	// } else if month == "march" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else if month == "april" {
	// 	fmt.Printf("%q has 30 days.\n", month)
	// } else if month == "may" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else if month == "june" {
	// 	fmt.Printf("%q has 30 days.\n", month)
	// } else if month == "july" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else if month == "august" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else if month == "september" {
	// 	fmt.Printf("%q has 30 days.\n", month)
	// } else if month == "october" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else if month == "november" {
	// 	fmt.Printf("%q has 30 days.\n", month)
	// } else if month == "december" {
	// 	fmt.Printf("%q has 31 days.\n", month)
	// } else {
	// 	fmt.Printf("%q is not a month.\n", month)
	// }

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name.")
		return
	}

	days := 28

	year := time.Now().Year()
	isLeapYear := year%400 == 0 || (year%100 != 0 && year%4 == 0)

	m := strings.ToLower(os.Args[1])
	if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "february" {
		if isLeapYear {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a valid month.\n", m)
		return
	}

	fmt.Printf("%q has %d days.\n", m, days)
}
