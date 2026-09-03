// Question 1.3 — Duplicate lines
//
// Read lines from standard input. Count how many times each distinct line
// appears. After EOF, print only the lines whose count is greater than 1,
// together with that count.
//
// You will need: bufio.Scanner, os.Stdin, map[string]int, ranging over a map.
//
// Run:
//   cat ../input/1.3-dup.input.txt | go run 1.3-dup.go
// Or type lines, then Ctrl+D (EOF on macOS/Linux).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	count := make(map[string]int)

	for scanner.Scan() {
		count[scanner.Text()]++
	}

	for line, n := range count {
		if n <= 1 {
			continue
		}
		fmt.Printf("%d\t%s\n", n, line)
	}
}
