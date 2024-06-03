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
	"os/exec"
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	sep := placeholder{
		" ",
		"░",
		" ",
		"░",
		" ",
	}

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		t := time.Now()

		// To print the hour and minute for the current time, split each number into its "10s" and
		// "1s" so we can get that digit out of the digits array, since the digits array only goes
		// from 0-9
		// Example: 10:55
		// 10 is "1" and "0"
		// 55 is "5" and "5"
		//
		// We can split the 2-digit number by first dividing it by 10 to get the 10s, and then
		// doing a modulo to get the 1s (the remainder).
		// We do this for both the minute and the hour.
		clock := [...]placeholder{
			digits[t.Hour()/10],
			digits[t.Hour()%10],
			sep,
			digits[t.Minute()/10],
			digits[t.Minute()%10],
		}

		// Print the time on the clock.
		// Each digit has 5 lines (elements), so we need to print each line of each digit across
		// the terminal before going to the next line.
		for line := range 5 {
			for digit := range clock {
				fmt.Printf("%s  ", clock[digit][line])
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
