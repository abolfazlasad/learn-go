# Chapter 8 — Goroutines and Channels

**Goal:** concurrency with goroutines and channels. Prefer communicating over sharing (sharing is Chapter 9).

Always have a receive for every send you need, or you will deadlock. Use `go run` and Ctrl+C to stop servers.

Write each section as its own file in this folder. Run one file at a time: `go run 8.1-goroutines.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [8.1 Goroutines](#81-goroutines)
- [8.2 Example: Concurrent Clock Server](#82-example-concurrent-clock-server)
- [8.3 Example: Concurrent Echo Server](#83-example-concurrent-echo-server)
- [8.4 Channels](#84-channels)
- [8.5 Looping in Parallel](#85-looping-in-parallel)
- [8.6 Example: Concurrent Web Crawler](#86-example-concurrent-web-crawler)
- [8.7 Multiplexing with select](#87-multiplexing-with-select)
- [8.8 Example: Concurrent Directory Traversal](#88-example-concurrent-directory-traversal)
- [8.9 Cancellation](#89-cancellation)
- [8.10 Example: Chat Server](#810-example-chat-server)

---

## 8.1 Goroutines

**File:** `8.1-goroutines.go`

**Learn:** `go f()`, independent execution, main exiting kills goroutines

- [ ] `go spinner()` (print a rotating `|/-\\`) while `main` computes something slow (`fib(45)` is enough).
- [ ] Start a goroutine that sleeps then prints. Let `main` return immediately — the print may never appear. Then wait (sleep or channel) so it does.
- [ ] **Problem:** launch 5 goroutines that each print their id. `main` waits using a channel of tokens (not `WaitGroup` yet unless you already know it).

---

## 8.2 Example: Concurrent Clock Server

**File:** `8.2-example-concurrent-clock-server.go`

**Learn:** `net.Listen`, `Accept` loop, one goroutine per connection, `net.Conn`

- [ ] Sequential clock: accept, write time once per second, close. Second client waits.
- [ ] `go handleConn(c)` so two `nc localhost 8000` sessions both tick.
- [ ] **Problem:** port from a flag. Timezone from env `TZ` or a flag (e.g. `US/Eastern`). Run two clocks on two ports.
- [ ] **Problem:** a client that dials several clock servers and prints a table of times (even if you stagger `fmt` roughly).

---

## 8.3 Example: Concurrent Echo Server

**File:** `8.3-example-concurrent-echo-server.go`

**Learn:** `bufio` + conn, echo, reverb (delayed copies), independent goroutines per conn

- [ ] Echo each line back. One goroutine per connection.
- [ ] For each line, launch goroutines that echo it after 0s, 1s, 2s with a shrinking prefix (reverb).
- [ ] **Problem:** if the client closes, in-flight reverb goroutines should stop or at least not crash the server. Defer `conn.Close()`.
- [ ] **Problem:** handle TCP vs a simple `net.Pipe` test in `main` if you do not want to use `nc`.

---

## 8.4 Channels

**File:** `8.4-channels.go`

**Learn:** `make(chan T)`, send/receive, close, range, unbuffered vs buffered, direction `chan<-` / `<-chan`

- [ ] Unbuffered ping-pong between two goroutines. `main` waits for “done”.
- [ ] Buffered `make(chan int, 3)`: send 3 without a receiver, then receive.
- [ ] Close a channel; `range ch` in the consumer. Send after close — panic, then don't.
- [ ] Function `func squares(out chan<- int)` vs receive-only parameter.
- [ ] **Problem:** pipeline: generator → square → print. Close so `range` ends.
- [ ] **Problem:** `countdown` that sends 3,2,1 then closes. Receiver prints `go!` after range ends.

---

## 8.5 Looping in Parallel

**File:** `8.5-looping-in-parallel.go`

**Learn:** bounded parallelism, `sync.WaitGroup` or counting semaphore via buffered channel

- [ ] Sequential `md5` (or just `len` + sleep) of several files.
- [ ] Unbounded: one goroutine per file. Collect results on a channel.
- [ ] Bounded: semaphore `make(chan struct{}, 2)` so at most 2 run at once.
- [ ] **Problem:** directory of files (or hardcoded paths). Print name + fake checksum. Faster than sequential, not 100 goroutines for 100 files.
- [ ] **Problem:** first error cancels starting new work (full cancel is 8.9). At least stop launching.

---

## 8.6 Example: Concurrent Web Crawler

**File:** `8.6-example-concurrent-web-crawler.go`

**Learn:** work queue, seen map, concurrency, avoiding unbounded crawl

- [ ] Sequential crawler: fetch URL, extract links (regex or `golang.org/x/net/html` if you add the module), enqueue unseen.
- [ ] Concurrent: tokens / semaphore + `seen` map. Protect `seen` with a mutex or give it to one goroutine (mutex is Chapter 9; a single “seen” goroutine is in the spirit of this chapter).
- [ ] **Problem:** depth limit (`-depth 2`).
- [ ] **Problem:** crawl `https://go.dev` (or a local test HTTP server you write) and print unique URLs. Be polite: bound concurrency to ~5.

---

## 8.7 Multiplexing with select

**File:** `8.7-multiplexing-with-select.go`

**Learn:** `select`, default, timeout, `time.After`, nil channels to disable a case

- [ ] `select` between two channels. Send to one from another goroutine.
- [ ] `time.After(1 * time.Second)` as a timeout case.
- [ ] `default` for a non-blocking send/receive.
- [ ] Set a channel variable to `nil` to disable its case.
- [ ] **Problem:** countdown that also aborts if stdin gets a line (`bufio.Scanner` in a goroutine sending on `abort`).
- [ ] **Problem:** merge two `<-chan int` into one using a loop + `select`. Exit when both are closed.

---

## 8.8 Example: Concurrent Directory Traversal

**File:** `8.8-example-concurrent-directory-traversal.go`

**Learn:** walking a tree with goroutines, counting bytes, limiting disk parallelism

- [ ] Sequential `du`: print total size of a directory.
- [ ] Concurrent walk: each directory in a goroutine. Semaphore on directory reads.
- [ ] Print running totals periodically with `select` + `time.Ticker` (or a ticker goroutine).
- [ ] **Problem:** `./du ~` (or a smaller dir) prints final bytes and file count. Compare to `du -sh` roughly.
- [ ] **Problem:** `-v` verbose: print progress every 500ms.

---

## 8.9 Cancellation

**File:** `8.9-cancellation.go`

**Learn:** done channel, broadcasting cancel by close, `context` as the modern form (optional extra)

- [ ] `var done = make(chan struct{})`. Goroutines `select` on `<-done`. Close `done` from `main` on stdin line.
- [ ] Show that after close, every receiver unblocks.
- [ ] **Problem:** add cancel to your `du` or crawler: Ctrl+C or a timeout stops launching work. Use `signal.Notify` or `time.After`.
- [ ] **Problem (extra):** same pattern with `context.WithCancel` / `WithTimeout`.

---

## 8.10 Example: Chat Server

**File:** `8.10-example-chat-server.go`

**Learn:** broadcaster goroutine, incoming/outgoing channels, client set, avoiding stuck sends

- [ ] TCP chat: each client has `ch chan string`. A hub goroutine owns the client map.
- [ ] Messages from one client go to all others. Arrivals/departures announced.
- [ ] Outgoing send must not block the hub forever (`select` + default, or buffered per client, or drop).
- [ ] **Problem:** two `nc` sessions can talk. Third joins and sees a welcome.
- [ ] **Problem:** idle timeout disconnects a client (`select` on `time.After` in the read loop, or a deadline on the conn).
