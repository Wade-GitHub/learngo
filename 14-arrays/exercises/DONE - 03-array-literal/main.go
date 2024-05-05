// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	names := [3]string{"John", "Jane", "Dave"}
	distances := [5]int{1, 2, 3, 4, 5}
	dataBuffer := [5]byte{0x01, 0x02, 0x03, 0x04, 0x05}
	ratios := [1]float64{1.04}
	alives := [4]bool{true, true, false, false}
	// zero := [0]byte{}

	fmt.Println("names")
	fmt.Println("===================")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println("===================")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println()

	fmt.Println("dataBuffer")
	fmt.Println("===================")
	for i := 0; i < len(dataBuffer); i++ {
		fmt.Printf("dataBuffer[%d]: %d\n", i, dataBuffer[i])
	}
	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("===================")
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println("===================")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println()

	fmt.Println()

	// ===== For ranges =====
	fmt.Println("names")
	fmt.Println("===================")
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println("===================")
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}
	fmt.Println()

	fmt.Println("dataBuffer")
	fmt.Println("===================")
	for i, v := range dataBuffer {
		fmt.Printf("dataBuffer[%d]: %d\n", i, v)
	}
	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("===================")
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println("===================")
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}
	fmt.Println()
}
