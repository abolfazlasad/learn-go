// Question 8.1 — Goroutines
//
// Start a goroutine that prints a spinner (| / - \) while main computes something
// slow (a naive fib(40) is enough). When main finishes, print the result.
//
// Extra: launch 5 goroutines that each send their id on a channel. Main receives 5 times.
//
// You will need: go, time.Sleep, maybe chan.
//
// Run:
//   go run 8.1-spinner.go

package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}

func spinner(done chan bool) {
	chars := []rune{'|', '/', '-', '\\'}
	var i int

	for {
		select {
		case <-done:
			fmt.Printf("\r")
			return
		default:
			fmt.Printf("\r%c", chars[i%4])
			i++
			time.Sleep(time.Millisecond * 100)
		}

	}
}

func main() {
	// TODO: write your solution
	done := make(chan bool)

	go spinner(done)

	result := fib(43)
	done <- true
	fmt.Printf("fib result: %v\n", result)

	fmt.Println("--------------------------------------------------")

	taskCh := make(chan int)
	n := 5
	for i := range n {
		go func() {
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)))
			fmt.Printf("finished: %d >>>>\n", i)
			taskCh <- i
			fmt.Printf("finished: %d <<<<\n", i)
		}()
	}
	time.Sleep(time.Second * 2)
	for range n {
		fmt.Printf("%v\n", (<-taskCh))
	}
	time.Sleep(time.Second)
}
