// Question 2.5 — Named types
//
// type Celsius float64
// type Fahrenheit float64
//
// Write CToF and FToC. You must convert explicitly; assigning a Celsius to a
// Fahrenheit without conversion must not compile (try it, then fix it).
// Print freezing (0C) and boiling (100C) in both units.
//
// You will need: type, conversion T(v).
//
// Run:
//   go run 2.5-celsius.go

package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	// TODO: write your solution
	var c_freezing Celsius = 0
	var c_boiling Celsius = 100
	var f_freezing = CToF(c_freezing)
	var f_boiling = CToF(c_boiling)

	fmt.Printf("%1.0fC = %1.0fF\n", c_freezing, f_freezing)
	fmt.Printf("%1.0fC = %1.0fF\n", c_boiling, f_boiling)

}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
