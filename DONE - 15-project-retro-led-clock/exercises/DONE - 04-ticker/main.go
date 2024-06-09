// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Ticker: Slide the Clock
//
//  Your goal is slide the placeholders every second.
//  Please run the solution to see it in action.
//
//
//  THIS IS A HARD EXERCISE:
//  + It will take you days but it will be worth it.
//  + For experienced developers, this can take an hour or so.
//
//
//  1. You need to determine the starting and the ending digits to create
//     the sliding effect.
//
//
//  2. Each second, start from the next placeholder, skip the previous one.
//     This means: Only draw the next placeholders.
//
//     Like this:
//
//     12:40:31
//     2:40:31
//     40:31
//     0:31
//     :31
//     31
//     1
//
//
//  3. After the last placeholder is displayed, fill the lines for the missing
//     placeholders, and then start from the first placeholder. Draw it to the
//     right part of the screen.
//
//     Like this:
//
//     12:40:31
//     2:40:31
//     40:31
//     0:31
//     :31
//     31
//     1
//            1
//           12
//          12:
//         12:4
//        12:40
//       12:40:
//      12:40:3
//     12:40:31
//
//     As you can see, you need to draw the clock from the right part of the
//     screen, beginning from the first placeholder.
//
//
// HINTS
//   + You would need to clear the screen inside the loop instead of once.
//     Otherwise the previous placeholders will be left on the screen.
//
//
// EXPECTED OUTPUT
//   Please run the solution to see it in action. Do not look at the source-code
//   though.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	clockDigitsLength = 8
	timeToSleep       = time.Second
)

func main() {
	// screen.Clear()

	for {
		// screen.MoveTopLeft()

		// Clear the screen on the first round of the infinite loop.
		clearCmd := exec.Command("clear")
		clearCmd.Stdout = os.Stdout
		clearCmd.Run()

		// First loop: clock is displayed as full, then scrolls to the left, disappearing.
		// We need an outer loop that starts with the entire length of the clock, and the inner
		// loops can use to print smaller clocks on each iteration.
		for i := 0; i < clockDigitsLength; i++ {
			// Get the current time
			now := time.Now()
			hour, min, sec := now.Hour(), now.Minute(), now.Second()

			// Build the clock array
			clock := [clockDigitsLength]placeholder{
				digits[hour/10], digits[hour%10],
				colon,
				digits[min/10], digits[min%10],
				colon,
				digits[sec/10], digits[sec%10],
			}

			// Clear the screen again, ready to print the next clock digits
			clearCmd := exec.Command("clear")
			clearCmd.Stdout = os.Stdout
			clearCmd.Run()

			// A tricky set of loops:
			// We cannot simply just print each digit, because each digit is made up of 5
			// strings as an array. We need to print each line of each digit, going across all
			// the digits.
			// This is why we start the outer loop as "for line := range clock[0]". We're
			// looping over the first line of each digit, then the second line, and so on until
			// the last line.
			// Each digit has the same length, so we can use clock[0] as the range.
			for line := range clock[0] {
				// The number of digits we print to the terminal need to get smaller and
				// smaller to make the illusion of the clock scrolling left.
				// On the first iteration, this will be the whole 8 digits of the clock.
				// On the second iteration, this will be 7 digits of the clock.
				// And so on, until no digits are left to print.
				// This is where we use "i" from the outer-most loop.
				for index, digit := range clock[i:] {
					// "next" gets the line of each digit as we're iterating (remember: we're
					// getting the lines across the digits, not the digits themselves.)
					next := clock[i:][index][line]
					// We want the clock colon to blink every 2nd second.
					if digit == colon && sec%2 == 0 {
						next = "   "
					}
					fmt.Print(next, "  ")
				}
				// Print a new line so we can print the next line of digits.
				fmt.Println()
			}
			time.Sleep(timeToSleep)
		}

		// Second loop: clock is displayed empty, then scrolls in from the right to the left,
		// appearing as it goes along.
		// We need an outer loop that starts with the entire length of the clock, and the inner
		// loops can use to print larger clocks on each iteration.
		for i := 0; i <= clockDigitsLength; i++ {
			// Get the current time.
			now := time.Now()
			hour, min, sec := now.Hour(), now.Minute(), now.Second()

			// Build the clock.
			clock := [clockDigitsLength]placeholder{
				digits[hour/10], digits[hour%10],
				colon,
				digits[min/10], digits[min%10],
				colon,
				digits[sec/10], digits[sec%10],
			}

			// Clear the terminal for each iteration.
			clearCmd := exec.Command("clear")
			clearCmd.Stdout = os.Stdout
			clearCmd.Run()

			// Same as in the first loop for the disappearing clock: we want to iterate across
			// the lines of each digit.
			for line := range clock[0] {
				// Here we implement things slightly differently:
				// We build a string with the line values for each digit going accross, then
				// print the entire string, rather than printing each line individually.
				// This is because we have to scroll in from the right, which needs to use
				// a right-aligning Printf().
				lineToPrint := ""
				// Loop through each digit of the clock. The number of digits will get larger
				// as the outer-most loop iterates, therefore printing more digits.
				for index, digit := range clock[0:i] {
					// Like in the disappearing clock, we want to print the colon every 2nd
					// second.
					if digit == colon && sec%2 == 0 {
						lineToPrint = fmt.Sprintf("%s  %s", lineToPrint, "  ")
					} else {
						lineToPrint = fmt.Sprintf("%s  %s", lineToPrint, clock[index][line])
					}
				}
				// Print the lines of all the digits going across the clock, right-aligned to
				// 40 spaces. This is the length of the clock (each digit has 5 characters per
				// line, with 8 digits. 5 * 8 = 40).
				fmt.Printf("%40s", lineToPrint)
				fmt.Println()
			}
			time.Sleep(timeToSleep)
		}
	}

}
