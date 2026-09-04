// Question 8.4 — Channels
//
// Build a pipeline: a generator sends 1..10 on a channel and closes it.
// A squarer receives, sends n*n, closes its out channel.
// Main ranges over squares and prints them.
//
// Use chan<- and <-chan in the helper signatures if you can.
//
// You will need: make(chan T), close, range, send, receive.
//
// Run:
//   go run 8.4-pipeline.go

package main

import (
	"fmt"
	"time"
)

func generator(n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		defer time.Sleep(time.Second * 2)

		for i := range n {
			out <- i + 1
		}
	}()

	return out
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func main() {
	numbers := generator(10)
	squares := squarer(numbers)

	for i := range squares {
		fmt.Println(i)
	}
	fmt.Println("finished")
}
