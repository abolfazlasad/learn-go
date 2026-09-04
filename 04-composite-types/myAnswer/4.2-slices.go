// Question 4.2 — Slices
//
// Start from s := []int{0, 1, 2, 3, 4, 5}.
// Print s[1:4], len, and cap after make([]int, 0, 5) and after several appends.
//
// Write reverse(s []int) that reverses in place.
// Show that append's return value matters: ignore it once, then use it correctly.
//
// You will need: slice, make, append, len, cap.
//
// Run:
//   go run 4.2-slices.go

package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	array := [6]int{0, 1, 2, 3, 4, 5}
	s := array[2:5]
	fmt.Printf("s: %T\n", s)
	fmt.Printf("array: %T\n", array)
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("len(array): %v\n", len(array))
	fmt.Printf("address of underlying array in s: %p\n", s)
	fmt.Printf("address of array variable: %p\n", &array)
	fmt.Printf("cap(s): %v\n", cap(s))
	fmt.Printf("cap(array): %v\n", cap(array))

	fmt.Printf("--------------------------------\n")

	s2 := append(s, 200)
	s2[0] = 1000
	fmt.Printf("array: %v\n", array)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("cap(s): %v\n", cap(s))
	fmt.Printf("cap(s2): %v\n", cap(s2))
	fmt.Printf("cap(array): %v\n", cap(array))
	fmt.Printf("address of underlying array in s: %p\n", s)
	fmt.Printf("address of underlying array in s2: %p\n", s2)
	fmt.Printf("address of array variable: %p\n", &array)

	fmt.Printf("--------------------------------\n")
	reverse(array[:])
	fmt.Printf("array: %v\n", array)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s2: %v\n", s2)

	fmt.Printf("--------------------------------\n")
	reverse(s)
	fmt.Printf("array: %v\n", array)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s2: %v\n", s2)
}
