# Chapter 1 — Tutorial

**Goal:** ship small working programs on day one: CLI tools, HTTP, a bit of concurrency.

Write each section as its own file in this folder. Run one file at a time: `go run 1.1-hello-world.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [1.1 Hello, World](#11-hello-world)
- [1.2 Command-Line Arguments](#12-command-line-arguments)
- [1.3 Finding Duplicate Lines](#13-finding-duplicate-lines)
- [1.4 Animated GIFs](#14-animated-gifs)
- [1.5 Fetching a URL](#15-fetching-a-url)
- [1.6 Fetching URLs Concurrently](#16-fetching-urls-concurrently)
- [1.7 A Web Server](#17-a-web-server)
- [1.8 Loose Ends](#18-loose-ends)

---

## 1.1 Hello, World

**File:** `1.1-hello-world.go`

**Learn:** `package main`, `func main`, `import`, `fmt`, `go run` vs `go build`

- [ ] Write `1.1-hello-world.go` that prints a greeting with `fmt.Println`.
- [ ] Run it with `go run 1.1-hello-world.go` and with `go build -o hello 1.1-hello-world.go` then `./hello`.
- [ ] Change the message, rebuild, confirm the binary changed.
- [ ] Break it: drop `package main`, drop the quotes, typo `Println`. Read each error. Fix.
- [ ] **Problem:** print three lines — a greeting, a hardcoded date string, and `2+2=` followed by the computed result (not the character `4` inside a string).

---

## 1.2 Command-Line Arguments

**File:** `1.2-command-line-arguments.go`

**Learn:** `os.Args`, `for`, `range`, `len`, slicing, `strings.Join`

- [ ] Print `os.Args` as one value, then print `os.Args[0]` vs the rest. Comment what `[0]` is.
- [ ] Loop with a C-style `for i := 0; i < len(os.Args); i++` and print `i` plus the arg.
- [ ] Rewrite using `for i, arg := range os.Args[1:]`.
- [ ] **Problem: echo** — join remaining args with a space on one line. Write a `+=` version, then a `strings.Join` version.
- [ ] **Problem:** with no args, print usage to stderr and `os.Exit(1)`. If the first arg is `-n`, omit the trailing newline.

---

## 1.3 Finding Duplicate Lines

**File:** `1.3-finding-duplicate-lines.go`

**Learn:** `map`, `bufio.Scanner`, `os.Open`, `os.Stdin`, ranging over maps

- [ ] Read stdin line by line. Count how often each line appears in a `map[string]int`.
- [ ] After EOF, print only lines whose count is greater than 1.
- [ ] Accept filenames as args. Open each file. Count across all files. Close files.
- [ ] If a file fails to open, print the error to stderr and continue (do not crash).
- [ ] **Problem:** also print which filenames contained each duplicate line.
- [ ] **Problem:** stream very long input without storing every unique line forever — or document why you still need the map and what that costs.

---

## 1.4 Animated GIFs

**File:** `1.4-animated-gifs.go`

**Learn:** packages beyond `fmt`, `image`, writing binary to stdout, loops that generate data

- [ ] Write a program that emits a tiny GIF (even a 16×16 solid color is enough) to stdout. Redirect to `out.gif` and open it.
- [ ] Animate at least 8 frames. Change color or position each frame.
- [ ] Drive size, frames, or delay from command-line flags (`flag` package, or raw `os.Args`).
- [ ] **Problem:** bouncing square — a filled rectangle that reverses direction at the edges. Save `bounce.gif`.

---

## 1.5 Fetching a URL

**File:** `1.5-fetching-a-url.go`

**Learn:** `net/http`, `http.Get`, `resp.Body`, `io.Copy`, `defer`, status codes

- [ ] Fetch a URL from `os.Args[1]`. Copy the body to stdout. Close the body.
- [ ] Print `resp.Status` as well.
- [ ] If `http.Get` fails, print the error and exit non-zero. Do not ignore errors.
- [ ] If the URL has no scheme, prepend `https://` and retry once.
- [ ] **Problem:** fetch and print only the first 200 bytes of the body, then the total byte count.

---

## 1.6 Fetching URLs Concurrently

**File:** `1.6-fetching-urls-concurrently.go`

**Learn:** `go` statement, channels, `time`, waiting for results (preview of Chapter 8)

- [ ] Fetch several URLs one after another. Time the whole run.
- [ ] Fetch the same URLs concurrently. Each goroutine sends a summary string on a channel. Main prints them.
- [ ] Main must wait until every fetch finishes (receive once per URL).
- [ ] **Problem:** print elapsed time per URL and total wall time. Concurrent should be closer to the slowest URL than to the sum.
- [ ] **Problem:** if a fetch fails, still report the error on the channel. Do not deadlock.

---

## 1.7 A Web Server

**File:** `1.7-a-web-server.go`

**Learn:** `http.HandleFunc`, `http.ListenAndServe`, `ResponseWriter`, `Request`, query params

- [ ] Serve `/` with a handler that writes `hello`. Listen on `:8000`. Hit it in a browser or with `curl`.
- [ ] Add `/echo` that writes `r.URL.Path`.
- [ ] Add `/time` that writes the current time.
- [ ] **Problem:** `/count` returns how many requests the process has seen (simple counter is fine; races come in Chapter 9).
- [ ] **Problem:** `/lissajous` (or `/bounce`) returns a GIF from 1.4 as `image/gif`. Optional query `cycles` or `size`.

---

## 1.8 Loose Ends

**File:** `1.8-loose-ends.go`

**Learn:** `switch`, named types as a preview, zero values, comments, `godoc`/`go doc`

- [ ] Write a `switch` on an int: 1–7 print weekday names, default print `unknown`.
- [ ] Switch on a string command (`start`, `stop`, `status`) without `break` (Go breaks automatically). Show what happens if you add `fallthrough`.
- [ ] Print the zero value of `int`, `string`, `bool`, and a pointer using `%v` / `%#v`.
- [ ] Run `go doc fmt.Println` and `go doc net/http.Get`. Skim the signatures.
- [ ] **Problem:** a tiny CLI with `switch` on `os.Args[1]`: `help`, `version`, `sum` (sum remaining args as integers). Unknown command → stderr + exit 1.
