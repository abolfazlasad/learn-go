// Question 3.6 — Constants and iota
//
// Define KB, MB, GB, TB as successive powers of 1024 using iota.
// Print them.
//
// Extra: type Day int with Sunday = iota ... Saturday, and weekend(d Day) bool.
//
// You will need: const, iota, typed constants.
//
// Run:
//   go run 3.6-iota.go

package main

import "fmt"

const (
	KB = 1 << (10 * iota + 10)
	MB
	GB
	TB
)

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// TODO: write your solution
	fmt.Printf("KB: %v\n", KB)
	fmt.Printf("MB: %v\n", MB)
	fmt.Printf("GB: %v\n", GB)
	fmt.Printf("TB: %v\n", TB)

	fmt.Printf("--------------------------------\n")
	fmt.Printf("Sunday: %v\n", Sunday)
	fmt.Printf("Monday: %v\n", Monday)
	fmt.Printf("Tuesday: %v\n", Tuesday)
	fmt.Printf("Wednesday: %v\n", Wednesday)
	fmt.Printf("Thursday: %v\n", Thursday)
	fmt.Printf("Friday: %v\n", Friday)
	fmt.Printf("Saturday: %v\n", Saturday)
}
