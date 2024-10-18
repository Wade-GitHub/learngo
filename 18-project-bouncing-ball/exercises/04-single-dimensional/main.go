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
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

// ---------------------------------------------------------
// EXERCISE: Single Dimensional
//
//  In this exercise you will understand why I use
//  a multi-dimensional board slice instead of a
//  single-dimensional one.
//
//  1. Remove this:
//     board := make([][]bool, width)
//
//  2. Use this:
//     board := make([]bool, width*height)
//
//  3. Adjust the rest of the operations in the code to work
//     with this single-dimensional slice.
//
//     You'll see how hard it becomes to work with it.
//
// ---------------------------------------------------------

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		// initial velocities
		// ivx, ivy = 5, 2
		ivx, ivy = 1, 1
	)

	var (
		px, py int // ball position
		// ppx, ppy int        // previous ball position
		p, pp  int
		vx, vy = ivx, ivy // velocities

		cell rune // current cell (for caching)
	)

	// you can get the width and height using the screen package easily:
	// width, height := screen.Size()
	width, height := 20, 20

	// get the rune width of the ball emoji
	ballWidth := runewidth.RuneWidth(cellBall)

	// adjust the width and height
	width /= ballWidth
	height-- // there is a 1 pixel border in my terminal

	// create the board
	// board := make([][]bool, width)
	// for column := range board {
	// 	board[column] = make([]bool, height)
	// }

	// Create the board
	board := make([]bool, width*height)

	// [ _, _, _, _, _, | _, _, _, _, _, |  _,  _,  _,  _,  _ ]
	//   0, 1, 2, 3, 4, | 5, 6, 7, 8, 9, | 10, 11, 12, 13, 14

	// drawing buffer length
	// *2 for extra spaces
	// +1 for newlines
	bufLen := (width*2 + 1) * height

	// create a drawing buffer
	buf := make([]rune, 0, bufLen)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy * width

		// when the ball hits a border reverse its direction
		// if px <= 0 || px >= width-ivx {
		if px <= 0 || px >= width-ivx {
			vx *= -1
		}
		// if py <= 0 || py >= height-ivy {
		if py < width || py >= len(board)-width {
			vy *= -1
		}

		// p = px
		p = py

		// remove the previous ball and put the new ball
		// board[px][py], board[ppx][ppy] = true, false
		board[p], board[pp] = true, false

		// save the previous positions
		// ppx, ppy = px, py
		pp = p

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		// for y := range board[0] {
		// 	for x := range board {
		// 		cell = cellEmpty
		// 		if board[x][y] {
		// 			cell = cellBall
		// 		}
		// 		buf = append(buf, cell, ' ')
		// 	}
		// 	buf = append(buf, '\n')
		// }
		for pos := range board {
			cell = cellEmpty
			if board[pos] {
				cell = cellBall
			}
			buf = append(buf, cell, ' ')

			// if pos > 0 && pos%width == 0 {
			if (pos+1)%width == 0 {
				buf = append(buf, '\n')
			}
		}

		// print the buffer
		screen.MoveTopLeft()
		fmt.Print(string(buf))
		// fmt.Print(buf)
		// fmt.Println()
		// fmt.Println(buf)

		// for j := 0; j <= 100; j++ {
		// 	fmt.Println(j, "%", width, "=", j%width)
		// }
		fmt.Println()
		fmt.Println("py", py)
		fmt.Println("width", width)
		fmt.Println("height", height)
		fmt.Println("width*height", width*height)

		// slow down the animation
		// time.Sleep(speed)
		time.Sleep(time.Second)
	}
}
