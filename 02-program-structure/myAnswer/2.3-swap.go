// Question 2.3 — Pointers and swap
//
// Write swap(a, b *int) that swaps the values the pointers refer to.
// In main, declare two ints, print them, call swap, print them again.
//
// Also print the zero value of an uninitialized int, string, and bool.
//
// You will need: var, :=, &, *, func.
//
// Run:
//   go run 2.3-swap.go

package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	var sampleInt int
	var sampleString string
	var sampleBool bool
	fmt.Println(sampleInt)
	fmt.Println(sampleBool)
	fmt.Println(sampleString)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
