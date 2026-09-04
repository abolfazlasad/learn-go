// Question 4.4 — Structs
//
// type Point struct { X, Y float64 }
// Write distance(p, q Point) float64 (euclidean).
//
// Extra: type Employee struct with an embedded Address. Set City via promotion
// (e.City) and via e.Address.City.
//
// You will need: struct literals, embedding, math.Hypot or sqrt.
//
// Run:
//   go run 4.4-point.go

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Address struct {
	City string
}

type Employee struct {
	Name string
	Address
}

func main() {
	a := Point{0, 0}
	b := Point{3, 4}
	fmt.Println(distance(a, b))

	var e Employee
	e.City = "Tehran"
	fmt.Println(e)
	e.Address.City = "Isfahan"
	fmt.Println(e)
}
