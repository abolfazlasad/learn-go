// Question 1.5 — Fetch a URL
//
// Take a URL from os.Args[1]. Fetch it with net/http. Copy the response body
// to stdout. Print the HTTP status as well.
//
// Close the body. If Get fails, print the error and exit non-zero.
// Do not ignore errors.
//
// Extra: if the argument has no scheme (no "http://" or "https://"), prepend
// "https://" and try once more.
//
// You will need: net/http, io.Copy or reading resp.Body, defer, os.
//
// Run:
//   go run 1.5-fetch.go https://example.com

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// TODO: write your solution
	argUrl := "https://example.com"
	if len(os.Args) > 1 {
		argUrl = os.Args[1]
	}

	resp, err := http.Get(argUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
