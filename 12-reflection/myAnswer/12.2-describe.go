// Question 12.2 — reflect.Type and reflect.Value
//
// describe(v any) prints the type, kind, and value of v.
// Call it on an int, a string, a *int, a []int, and a small struct.
//
// You will need: reflect.TypeOf, reflect.ValueOf, Kind, fmt.
//
// Run:
//   go run 12.2-describe.go

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func describe(v any) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)

	fmt.Println("Type:", t)
	fmt.Println("Kind:", value.Kind())
	fmt.Println("Value:", value)
	fmt.Println("--------------------------------------------------")
}

func main() {
	// TODO: write your solution
	describe(3)
	describe(3.14)
	describe("hi")

	var a int = 2
	describe(&a)

	b := [3]int{1, 2, 3}
	describe(b)
	c := b[:]
	describe(c)

	describe(Person{
		Name: "Abolfazl",
		Age:  25,
	})
}
