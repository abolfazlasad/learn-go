// Question 1.2 — Echo
//
// Print the command-line arguments after the program name, joined by a single
// space, on one line.
//
// os.Args[0] is the program path. The real arguments start at os.Args[1].
// Try a loop first. Then, if you want, the same result with strings.Join.
//
// If there are no arguments, print a short usage message to stderr and exit
// with status 1 (os.Exit).
//
// You will need: os.Args, a for loop or range, fmt, maybe strings.
//
// Run:
//   go run 1.2-echo.go hello go
//   go run 1.2-echo.go

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	fmt.Println(strings.Join(os.Args[1:], " "))
}
