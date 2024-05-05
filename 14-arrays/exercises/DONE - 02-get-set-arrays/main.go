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
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
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
