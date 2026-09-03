# Chapter 9 — Concurrency with Shared Variables

**Goal:** races, mutexes, memory visibility, `sync.Once`, the race detector, a concurrent cache.

Run interesting programs with `go run -race .` once you reach 9.6 — and earlier whenever you share memory.

Write each section as its own file in this folder. Run one file at a time: `go run 9.1-race-conditions.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [9.1 Race Conditions](#91-race-conditions)
- [9.2 Mutual Exclusion: sync.Mutex](#92-mutual-exclusion-syncmutex)
- [9.3 Read/Write Mutexes: sync.RWMutex](#93-readwrite-mutexes-syncrwmutex)
- [9.4 Memory Synchronization](#94-memory-synchronization)
- [9.5 Lazy Initialization: sync.Once](#95-lazy-initialization-synconce)
- [9.6 The Race Detector](#96-the-race-detector)
- [9.7 Example: Concurrent Non-Blocking Cache](#97-example-concurrent-non-blocking-cache)
- [9.8 Goroutines and Threads](#98-goroutines-and-threads)

---

## 9.1 Race Conditions

**File:** `9.1-race-conditions.go`

**Learn:** data race vs race condition, shared variables without sync

- [ ] Two goroutines increment the same `int` 100000 times each with no lock. Print the total (often not 200000).
- [ ] Same with a `map` write from two goroutines — often panics.
- [ ] **Problem:** bank-balance example: `deposit` and `withdraw` from two goroutines. Show a lost update.
- [ ] **Problem:** write a *correct* version using a single monitor goroutine and commands on a channel (no mutex yet). Compare to the broken one.

---

## 9.2 Mutual Exclusion: sync.Mutex

**File:** `9.2-mutual-exclusion-sync-mutex.go`

**Learn:** `Lock`/`Unlock`, `defer Unlock`, protecting invariants, don't copy mutexes

- [ ] Protect the counter from 9.1 with `sync.Mutex`. Total must be exact.
- [ ] `defer mu.Unlock()` after `Lock`.
- [ ] Embed `sync.Mutex` in a struct `Bank { mu sync.Mutex; balance int }`. Methods `Deposit`, `Balance`.
- [ ] **Problem:** `Withdraw(amount) bool` that fails if funds are insufficient. The check+subtract must be one critical section.
- [ ] **Problem:** document in a comment what the mutex guards (which fields).

---

## 9.3 Read/Write Mutexes: sync.RWMutex

**File:** `9.3-read-write-mutexes-sync-rwmutex.go`

**Learn:** many readers xor one writer

- [ ] Cache `map[string]string` with `RLock` on get and `Lock` on set.
- [ ] Hammer gets from many goroutines and rare sets. Compare to a plain `Mutex` (optional timing).
- [ ] **Problem:** do not hold `RLock` and then try `Lock` in the same goroutine (deadlock). Write a small program that deadlocks, then fix it.

---

## 9.4 Memory Synchronization

**File:** `9.4-memory-synchronization.go`

**Learn:** a mutex/channel happens-before; no “benign” races; visibility

- [ ] Writer sets `a = 1` then `b = 2` without sync; reader prints `b` then `a`. You may or may not see a surprise — still treat it as illegal.
- [ ] Same with mutex around both writes and both reads. Now the pairing is defined.
- [ ] **Problem:** one-page `notes.md`: in your own words, why `go run -race` matters even if the program “looks fine.”

---

## 9.5 Lazy Initialization: sync.Once

**File:** `9.5-lazy-initialization-sync-once.go`

**Learn:** `Once.Do`, init once even with concurrent callers

- [ ] `var icons map[string]image.Image` (or `map[string]string` to keep it simple) loaded in `loadIcons`. Call from many goroutines with `once.Do(loadIcons)`.
- [ ] Broken version: `if icons == nil { load() }` from many goroutines. Show the race with `-race`.
- [ ] **Problem:** `func Get(name string) string` that lazy-loads a config file once.

---

## 9.6 The Race Detector

**File:** `9.6-the-race-detector.go`

**Learn:** `go run -race`, `go test -race`, how to read the report

- [ ] Re-run the unsynchronized counter with `go run -race .`. Save the report in a comment or `race.txt`.
- [ ] Fix it. `-race` should be silent.
- [ ] **Problem:** pick any Chapter 8 program that shared a map. Run `-race`. Fix or confirm it is clean.

---

## 9.7 Example: Concurrent Non-Blocking Cache

**File:** `9.7-example-concurrent-non-blocking-cache.go`

**Learn:** memoizing a function, duplicate suppression, `sync.Mutex` + channels or `singleflight` idea by hand

- [ ] Sequential memo: `map[string]result` for an expensive `httpGetBody(url)` (or `time.Sleep` fake).
- [ ] Add a mutex. Concurrent callers still duplicate in-flight work — show that with prints.
- [ ] **Problem:** each URL maps to a `*entry` with a channel that is closed when ready. Callers wait on that channel. Duplicate URLs share one fetch.
- [ ] **Problem:** handle error results (cache failures or not — pick one and comment why).
- [ ] Run with `-race`.

---

## 9.8 Goroutines and Threads

**File:** `9.8-goroutines-and-threads.go`

**Learn:** M:N scheduler (concept), GOMAXPROCS, stacks grow, don't assume OS thread identity

- [ ] Print `runtime.GOMAXPROCS(0)` and `runtime.NumCPU()`. Set `GOMAXPROCS=1` and run a CPU-bound and an IO-bound tiny demo.
- [ ] `runtime.LockOSThread()` experiment: optional; if you skip, write `notes.md` instead.
- [ ] **Problem:** busy-loop goroutine vs sleeping goroutine. With `GOMAXPROCS=1`, does the sleeper still run? Prove with prints.
- [ ] **Problem:** `notes.md` — 10–15 lines: goroutine vs OS thread, what you can and cannot assume.
