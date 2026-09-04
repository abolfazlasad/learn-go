// Question 13.2 — unsafe.Pointer
//
// Read math.Float64bits in go doc or the stdlib. Then print the bits of 1.0
// using math.Float64bits (preferred).
//
// In a comment, explain why holding a uintptr across a function call is unsafe
// for the GC. Do not write a crashing example.
//
// You will need: math.Float64bits; maybe unsafe.Pointer only in comments.
//
// Run:
//   go run 13.2-pointer.go

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1.0

	bits := math.Float64bits(x)

	fmt.Printf("%064b\n", bits)
}
