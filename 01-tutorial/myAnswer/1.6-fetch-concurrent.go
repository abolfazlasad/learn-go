// Question 1.6 — Fetch URLs concurrently
//
// Fetch several URLs at the same time. Each goroutine sends a one-line summary
// on a channel (URL, status or error, elapsed time). Main receives one result
// per URL, then prints total wall time.
//
// Do not deadlock if a fetch fails: still send on the channel.
// Concurrent wall time should be closer to the slowest URL than to the sum.
//
// You will need: go, chan, time, net/http.
//
// Run:
//   go run 1.6-fetch-concurrent.go https://example.com https://go.dev

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		urls = []string{"https://example.com", "https://go.dev", "https://www.google.com"}
	}

	ch := make(chan string)
	var wg sync.WaitGroup
	start := time.Now()
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fetch(url, ch)
		}(url)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for line := range ch {
		fmt.Println(line)
	}
	fmt.Printf("%vms elapsed\n", time.Since(start).Milliseconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer resp.Body.Close()
	ch <- fmt.Sprintf("%vms %s", time.Since(start).Milliseconds(), url)
}
