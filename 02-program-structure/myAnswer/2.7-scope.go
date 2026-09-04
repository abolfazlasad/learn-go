// Question 2.7 — Scope
//
// In main, declare x := 1. Inside a nested block, declare x := 2 and print it.
// After the block, print x again. The outer x must still be 1.
//
// Then try to use a variable declared in an if short statement after the if.
// That must not compile. Leave a comment explaining why, and delete the bad line.
//
// You will need: blocks { }, :=, if v := ...
//
// Run:
//   go run 2.7-scope.go

package main

import "fmt"

func main() {
	// TODO: write your solution
	x := 1
	{
		x := 2
		fmt.Printf("x: %v\n", x)
	}
	fmt.Printf("x: %v\n", x)

	if x := -1; x < 0 {
		fmt.Printf("x: %v\n", x)
	}
	fmt.Printf("x: %v\n", x)

}
