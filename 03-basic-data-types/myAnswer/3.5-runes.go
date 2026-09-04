// Question 3.5 — Strings vs runes
//
// Let s := "سلام". Print:
//   len(s)                      // bytes
//   utf8.RuneCountInString(s)   // runes
// Then range over s with for i, r := range s and print i and r.
//
// Extra: basename(path string) string — strip the directory and an optional
// suffix. Example: basename("/tmp/a.go", ".go") ideas are in the book; write
// your own from scratch.
//
// You will need: string, range, unicode/utf8, maybe strings.
//
// Run:
//   go run 3.5-runes.go

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// TODO: write your solution
	input := "سلامali"
	len_s := len(input)
	rune_count := utf8.RuneCountInString("سلام")
	fmt.Printf("len_s: %v\n", len_s)
	fmt.Printf("rune_count: %v\n", rune_count)

	for i, r := range input {
		fmt.Printf("i: %v, r: %v\n", i, string(r))
	}
}
