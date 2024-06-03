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

const placeholder = "|"

func main() {
	type digit [5]string

	// "||    "
	left := fmt.Sprintf("%s%s", strings.Repeat(placeholder, 2), strings.Repeat(" ", 4))
	// "  ||  "
	// middle := fmt.Sprintf("%s%s%s", strings.Repeat(" ", 2), strings.Repeat(placeholder, 2), strings.Repeat(" ", 2))
	// "    ||"
	right := fmt.Sprintf("%s%s", strings.Repeat(" ", 4), strings.Repeat(placeholder, 2))
	// "||  ||"
	split := fmt.Sprintf("%s%s%s", strings.Repeat(placeholder, 2), strings.Repeat(" ", 2), strings.Repeat(placeholder, 2))
	// "||||||"
	full := strings.Repeat(placeholder, 6)

	// zero := [5]string{
	// 	"||||||",
	// 	"||  ||",
	// 	"||  ||",
	// 	"||  ||",
	// 	"||||||",
	// }

	zero := digit{
		full,
		split,
		split,
		split,
		full,
	}

	// one := [5]string{
	// 	"  ||  ",
	// 	"  ||  ",
	// 	"  ||  ",
	// 	"  ||  ",
	// 	"  ||  ",
	// }

	one := digit{
		right,
		right,
		right,
		right,
		right,
	}

	// two := [5]string{
	// 	"||||||",
	// 	"    ||",
	// 	"||||||",
	// 	"||    ",
	// 	"||||||",
	// }

	two := digit{
		full,
		right,
		full,
		left,
		full,
	}

	// three := [5]string{
	// 	"||||||",
	// 	"    ||",
	// 	"||||||",
	// 	"    ||",
	// 	"||||||",
	// }

	three := digit{
		full,
		right,
		full,
		right,
		full,
	}

	// four := [5]string{
	// 	"||  ||",
	// 	"||  ||",
	// 	"||||||",
	// 	"    ||",
	// 	"    ||",
	// }

	four := digit{
		split,
		split,
		full,
		right,
		right,
	}

	// five := [5]string{
	// 	"||||||",
	// 	"||    ",
	// 	"||||||",
	// 	"    ||",
	// 	"||||||",
	// }

	five := digit{
		full,
		left,
		full,
		right,
		full,
	}

	// six := [5]string{
	// 	"||||||",
	// 	"||    ",
	// 	"||||||",
	// 	"||  ||",
	// 	"||||||",
	// }

	six := digit{
		full,
		left,
		full,
		split,
		full,
	}

	// seven := [5]string{
	// 	"||||||",
	// 	"    ||",
	// 	"    ||",
	// 	"    ||",
	// 	"    ||",
	// }

	seven := digit{
		full,
		right,
		right,
		right,
		right,
	}

	// eight := [5]string{
	// 	"||||||",
	// 	"||  ||",
	// 	"||||||",
	// 	"||  ||",
	// 	"||||||",
	// }

	eight := digit{
		full,
		split,
		full,
		split,
		full,
	}

	// nine := [5]string{
	// 	"||||||",
	// 	"||  ||",
	// 	"||||||",
	// 	"    ||",
	// 	"    ||",
	// }

	nine := digit{
		full,
		split,
		full,
		right,
		right,
	}

	digits := [...]digit{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}

	// To print each digit side-by-side, we ned to print each line of each digit.
	// Each digit has 5 lines.
	// for i := 0; i < 5; i++ {
	for line := range 5 {
		for _, digit := range digits {
			// fmt.Printf("%s  ", digit[i])
			fmt.Printf("%s  ", digit[line])
		}
		fmt.Println()
	}
}
