// Question 5.8 — defer
//
// In one function, defer three fmt.Println calls with "first", "second", "third".
// Confirm they run LIFO (third, second, first) after the function body prints "body".
//
// Extra: timeit(name string, f func()) that defers a duration print using time.Now().
//
// You will need: defer, time.
//
// Run:
//   go run 5.8-defer.go

package main

import (
	"fmt"
	"time"
)

func f1() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")

	fmt.Println("body")
}

func timeit(name string, f func()) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("duration %s: %v\n", name, duration)
	}()
	f()
}

func main() {
	f1()
	timeit("f1", f1)

}
