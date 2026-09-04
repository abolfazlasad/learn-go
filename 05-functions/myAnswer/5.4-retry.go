// Question 5.4 — Errors
//
// Write a function that fails sometimes (for example fail if n is even).
// Write retry(f func() error, times int) error that calls f up to times.
// If all fail, return the last error.
//
// In main, never ignore err. Print it.
//
// You will need: error, errors.New, fmt.Errorf, err != nil.
//
// Run:
//   go run 5.4-retry.go

package main

import (
	"fmt"
)

func errorIfEven(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("even: n=%d is even", n)
	}
	return nil
}

func retry(f func() error, times int) error {
	var err error
	for range times {
		err = f()
		if err == nil {
			return nil
		}
	}
	return err
}

func main() {
	err := retry(func() error { return errorIfEven(2) }, 3)
	if err != nil {
		fmt.Println(err)
	}
}
