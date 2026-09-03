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

func main() {
	// TODO: write your solution
}
