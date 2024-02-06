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
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

const (
	username = "jack"
	password = "1888"

	usage         = "Usage: [username] [password]"
	accessDenied  = `Access denied for "%s"` + "\n"
	invalidPass   = `Invalid password for "%s"` + "\n"
	accessGranted = `Access granted to "%s"` + "\n"
)

func main() {
	if len(os.Args)-1 != 2 {
		fmt.Println(usage)
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u != username {
		fmt.Printf(accessDenied, u)
	} else if p != password {
		fmt.Printf(invalidPass, u)
	} else {
		fmt.Printf(accessGranted, u)
	}
}
