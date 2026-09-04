// Question 4.3 — Maps
//
// Count word frequencies from command-line arguments (or stdin).
// Look up a missing key with v, ok := m[k].
// delete one key. Print keys in sorted order (collect keys, sort.Strings).
//
// Extra: equal(a, b map[string]int) bool.
//
// You will need: map, make, comma-ok, delete, sort.
//
// Run:
//   go run 4.3-wordcount.go go go rust go
//   cat ../input/4.3-wordcount.input.txt | go run 4.3-wordcount.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	wordCount := make(map[string]int)
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")
		for _, word := range words {
			wordCount[word]++
		}
	}

	fmt.Printf("--------------------------------\n")
	fmt.Printf("wordCount: %v\n", wordCount)

	fmt.Printf("--------------------------------\n")
	delete(wordCount, "test")
	fmt.Printf("wordCount: %v\n", wordCount)

	wordCount["zoo"] = 2
	wordCount["ali"] = 1

	fmt.Printf("--------------------------------\n")
	keys := make([]string, 0, len(wordCount))

	for k := range wordCount {
		keys = append(keys, k)
	}
	fmt.Printf("keys: %v\n", keys)
	sort.Strings(keys)
	fmt.Printf("keys: %v\n", keys)

}
