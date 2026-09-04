// Question 13.1 — Sizeof, Alignof, Offsetof
//
// Print unsafe.Sizeof for int, int32, int64, string, and []int.
//
// Define type S struct { A bool; B float64; C int32 }
// Print Sizeof(S{}), Alignof, and Offsetof each field.
// Reorder fields to reduce padding and print the new size.
// Draw the layout in a comment (offsets + padding).
//
// You will need: unsafe.Sizeof, Alignof, Offsetof.
//
// Run:
//   go run 13.1-layout.go

package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	A bool
	B float64
	C int32
}

type SOptimized struct {
	B float64
	A bool
	C int32
}

func main() {
	// Basic sizes
	fmt.Println("int:   ", unsafe.Sizeof(int(0)))
	fmt.Println("int32: ", unsafe.Sizeof(int32(0)))
	fmt.Println("int64: ", unsafe.Sizeof(int64(0)))
	fmt.Println("string:", unsafe.Sizeof(""))
	fmt.Println("[]int: ", unsafe.Sizeof([]int{}))
	fmt.Println("[]int{1, 2, ..., 8}: ", unsafe.Sizeof([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println("[3]int: ", unsafe.Sizeof([3]int{}))

	// S layout
	fmt.Println("\nS:")
	fmt.Println("Size:  ", unsafe.Sizeof(S{}))
	fmt.Println("Align: ", unsafe.Alignof(S{}))
	fmt.Println("A offset:", unsafe.Offsetof(S{}.A))
	fmt.Println("B offset:", unsafe.Offsetof(S{}.B))
	fmt.Println("C offset:", unsafe.Offsetof(S{}.C))

	// Optimized layout
	fmt.Println("\nSOptimized:")
	fmt.Println("Size:  ", unsafe.Sizeof(SOptimized{}))
	fmt.Println("Align: ", unsafe.Alignof(SOptimized{}))
	fmt.Println("B offset:", unsafe.Offsetof(SOptimized{}.B))
	fmt.Println("C offset:", unsafe.Offsetof(SOptimized{}.C))
	fmt.Println("A offset:", unsafe.Offsetof(SOptimized{}.A))
}
