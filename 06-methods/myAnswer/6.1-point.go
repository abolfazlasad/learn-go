// Question 6.1 — Methods
//
// type Point struct { X, Y float64 }
// func (p Point) Distance(q Point) float64
//
// Call p.Distance(q) from main.
// Also write the same logic as a package function Distance(p, q Point) and
// compare the call sites.
//
// You will need: method receiver, func (t T).
//
// Run:
//   go run 6.1-point.go

package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance1(q Point) float64 {
	return Distance(p, q)
}

func (p *Point) Distance2(q Point) float64 {
	return Distance(*p, q)
}


func main() {
	a := Point{0, 0}
	b := Point{3, 4}

	fmt.Println((a).Distance1(b)) // Go dereferences ptr; value receiver still works
	fmt.Println((&a).Distance1(b)) // Go dereferences ptr; value receiver still works
	fmt.Println((a).Distance2(b)) // Go dereferences ptr; value receiver still works
	fmt.Println((&a).Distance2(b)) // Go dereferences ptr; value receiver still works

}
