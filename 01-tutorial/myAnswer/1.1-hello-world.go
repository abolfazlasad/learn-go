// Question 1.1 — Hello, World
//
// Write a small program that prints three lines:
//   1. a greeting that includes your name
//   2. a date string you type yourself (do not look up time.Now yet)
//   3. the text 2+2= followed by the computed result (2+2 in code, not the character 4 inside a string)
//
// You will need: package main, func main, import "fmt", fmt.Println.
//
// Run:
//   go run 1.1-hello-world.go
// Then also:
//   go build -o hello 1.1-hello-world.go && ./hello

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Abolazl")
	var now = time.Now()
	fmt.Printf("%v %v %v\n", now.Day(), now.Month(), now.Year())
	fmt.Printf("2+2=%d", 2+2)
}
