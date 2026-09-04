// Question 3.1 — Integers and bits
//
// Write popcount(n uint64) int: how many 1-bits in n.
// Print results for 0, 1, 255, and ^uint64(0) (all bits set).
//
// Extra: print a few numbers in binary with fmt "%08b" or "%b".
//
// You will need: uint64, bit ops (&, >>, or a loop), fmt.
//
// Run:
//   go run 3.1-popcount.go

package main

import (
	"fmt"
)

func popcount(n uint64) int {
	s := fmt.Sprintf("%b", n)
	// fmt.Printf("s: %v\n", s)

	count := 0
	for _, ch := range s {
		if ch == '1' {
			count += 1
		}
	}

	return count
}

func main() {
	fmt.Printf("popcount(0): %v\n", popcount(0))
	fmt.Printf("popcount(1): %v\n", popcount(1))
	fmt.Printf("popcount(5): %v\n", popcount(5))
	fmt.Printf("popcount(255): %v\n", popcount(255))
	fmt.Printf("popcount(^uint64(0)): %v\n", popcount(^uint64(0)))
}
