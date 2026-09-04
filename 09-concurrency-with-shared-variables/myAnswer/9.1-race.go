// Question 9.1 — Race conditions
//
// Two goroutines each increment the same int 100000 times with no lock.
// Print the total (often not 200000).
//
// Then run: go run -race 9.1-race.go
//
// Extra: the same bug with a map write from two goroutines (it may panic).
//
// You will need: go, shared variable, time.Sleep or sync.WaitGroup to wait.
//
// Run:
//   go run 9.1-race.go
//   go run -race 9.1-race.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()

		for range 100000 {
			counter++
		}
	}

	for range 2 {
		wg.Add(1)
		go increment()
	}

	wg.Wait()

	fmt.Printf("counter: %v\n", counter)
}
