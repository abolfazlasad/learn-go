# Chapter 5 — Functions

**Goal:** declarations, recursion, multiple returns, errors, function values, closures, variadic, defer, panic, recover.

Write each section as its own file in this folder. Run one file at a time: `go run 5.1-function-declarations.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [5.1 Function Declarations](#51-function-declarations)
- [5.2 Recursion](#52-recursion)
- [5.3 Multiple Return Values](#53-multiple-return-values)
- [5.4 Errors](#54-errors)
- [5.5 Function Values](#55-function-values)
- [5.6 Anonymous Functions](#56-anonymous-functions)
- [5.7 Variadic Functions](#57-variadic-functions)
- [5.8 Deferred Function Calls](#58-deferred-function-calls)
- [5.9 Panic](#59-panic)
- [5.10 Recover](#510-recover)

---

## 5.1 Function Declarations

**File:** `5.1-function-declarations.go`

**Learn:** parameters, results, named results, no default args, pass by value

- [ ] `func add(a int, b int) int` and the shortened `func add(a, b int) int`.
- [ ] Named result: `func split(sum int) (x, y int)`. Use a bare `return`.
- [ ] Pass a struct into a function, mutate a field, show the caller did not change. Then pass a pointer.
- [ ] **Problem:** `hypot(x, y float64) float64`. Call it. Try calling with one argument — must not compile.

---

## 5.2 Recursion

**File:** `5.2-recursion.go`

**Learn:** recursive functions, base case, stack growth

- [ ] `fact(n int) int`. Check 0, 1, 5, 10.
- [ ] Walk a small `map[string][]string` HTML-like tree (or a directory listing you hardcode) recursively and print names with indent.
- [ ] **Problem:** `walk(dir string)` using `os.ReadDir`. Print every file path. Skip errors on one entry, keep walking.
- [ ] **Problem:** Fibonacci both recursive and loop. Time n=40. Comment why the naive recursive one explodes.

---

## 5.3 Multiple Return Values

**File:** `5.3-multiple-return-values.go`

**Learn:** `(T, error)`, blank identifier `_`, returning extra results

- [ ] Function that returns `(int, error)` for parsing. Handle both results.
- [ ] Use `_` to ignore a result. Then stop ignoring errors.
- [ ] **Problem:** `minmax(nums []int) (min, max int, err error)` — error on empty slice.
- [ ] **Problem:** `div(a, b int) (q, r int, err error)`.

---

## 5.4 Errors

**File:** `5.4-errors.go`

**Learn:** `error` interface, `errors.New`, `fmt.Errorf`, checking `err != nil`, wrapping (light)

- [ ] Return `errors.New("empty")` and `fmt.Errorf("open %s: %v", name, err)`.
- [ ] Write a function that fails. In `main`, print `err` and exit. Never `_ = err` for real failures.
- [ ] **Problem:** read a file of integers (one per line). Collect parse errors; still return the numbers that worked, plus a combined error (or a slice of errors).
- [ ] **Problem:** retry a flaky function up to 3 times. If all fail, return the last error.

---

## 5.5 Function Values

**File:** `5.5-function-values.go`

**Learn:** functions as values, `func` types, passing callbacks, nil function values

- [ ] `var f func(int) int = square`. Call `f(3)`.
- [ ] `apply([]int, func(int) int) []int` that maps the function over the slice.
- [ ] Assign `f = nil` and call it — recover from the panic in a comment after you see it, then nil-check.
- [ ] **Problem:** `strings.Map`-style: `mapRunes(s string, f func(rune) rune) string`.
- [ ] **Problem:** sort a slice of strings by length using `sort.Slice` and a function value.

---

## 5.6 Anonymous Functions

**File:** `5.6-anonymous-functions.go`

**Learn:** closures, capturing loop variables (modern Go vs classic bug), `go func`

- [ ] `func() { fmt.Println("hi") }()` immediately invoked.
- [ ] `makeAdder(n int) func(int) int` that returns a closure. Call twice with different `n`.
- [ ] Launch several goroutines in a loop that print the loop index. First capture the variable so it is correct. (If you see the same index, fix it.)
- [ ] **Problem:** `once(f func()) func()` — returned function runs `f` only the first time.
- [ ] **Problem:** topological-style visit of a graph with an inner recursive function `var visit func(string)`.

---

## 5.7 Variadic Functions

**File:** `5.7-variadic-functions.go`

**Learn:** `...T`, passing a slice with `s...`

- [ ] `sum(nums ...int) int`. Call as `sum(1,2,3)` and `sum(s...)`.
- [ ] `fmt.Println` is variadic — wrap it: `logf(format string, args ...any)`.
- [ ] **Problem:** `max(first int, rest ...int) int`. Require at least one argument (so empty is a compile error, not a runtime one).
- [ ] **Problem:** `join(sep string, parts ...string) string` without calling `strings.Join` (you may use a `Builder`).

---

## 5.8 Deferred Function Calls

**File:** `5.8-deferred-function-calls.go`

**Learn:** `defer` LIFO, args evaluated immediately, `defer` for Close, named results + defer

- [ ] Three `defer fmt.Println(...)` in one function. Confirm LIFO order.
- [ ] `defer` a function with an argument. Change the variable after `defer`. Print: the arg was already evaluated.
- [ ] Open a file, `defer f.Close()`, read it. Check Close's error in a named-result function if you can.
- [ ] **Problem:** `timeit(name string, f func())` using `defer` and `time.Now()` to print duration even if `f` panics (you can recover later; here just defer the print).
- [ ] **Problem:** write a function that defers unlocking a mutex (`sync.Mutex`) around a map update.

---

## 5.9 Panic

**File:** `5.9-panic.go`

**Learn:** panic for truly impossible cases, stack unwind, vs errors for expected failure

- [ ] `panic("boom")` from a helper. Read the stack trace.
- [ ] Panic on a nil pointer dereference (`var p *int; fmt.Println(*p)`).
- [ ] **Problem:** `mustParse(s string) int` that panics if parse fails. Call it from `main` with good and bad input (bad input on purpose, once).
- [ ] **Problem:** rewrite `mustParse` as `parse(s string) (int, error)`. Use panic only in `main` if you choose to abort. Comment when you would actually panic in library code.

---

## 5.10 Recover

**File:** `5.10-recover.go`

**Learn:** `recover` only inside deferred functions, converting panic to error

- [ ] `defer func() { if r := recover(); r != nil { ... } }()` around a panic. Print `r`.
- [ ] Call `recover()` *not* in a deferred function — show it does nothing useful.
- [ ] **Problem:** `safeCall(f func()) (err error)` that turns a panic into `fmt.Errorf("panic: %v", r)`.
- [ ] **Problem:** a worker that panics on one item must not kill the loop — recover per item, continue.
