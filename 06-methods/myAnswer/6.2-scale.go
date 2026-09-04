// Question 6.2 — Pointer receivers
//
// func (p *Point) ScaleBy(factor float64) that mutates p.
// Show that a value-receiver version would not persist the change.
//
// type Counter struct { n int } with Inc() (pointer) and Value() int.
//
// You will need: func (p *T), addressable values.
//
// Run:
//   go run 6.2-scale.go

package main

import "fmt"

type Point struct{ X, Y float64 }
type Counter struct{ n int }

func (p *Point) ScaleBy1(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleBy2(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func (p *Point) print() {
	fmt.Printf("Point: %v\n", *p)
}

func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Value() int {
	return c.n
}

func main() {
	a := Point{1, 2}
	a.print()
	a.ScaleBy1(2)
	a.print()
	a.ScaleBy2(2)
	a.print()

	c := Counter{0}
	c.Inc()
	fmt.Printf("Counter: %v\n", c.Value())
	fmt.Printf("c.n: %v\n", c.n)
}
