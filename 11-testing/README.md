# Chapter 11 — Testing

**Goal:** `go test` is how you know the code is true. Write tests first when a problem is crisp; write them immediately after when you are exploring.

Put library code and `_test.go` in this chapter folder, named with the section prefix (`11.2-word.go`, `11.2-word_test.go`). Extract a small package only when `go test` cannot share a file with another section's `main`.

Write each section as its own file in this folder. Run one file at a time: `go run 11.1-the-go-test-tool.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [11.1 The go test Tool](#111-the-go-test-tool)
- [11.2 Test Functions](#112-test-functions)
- [11.3 Coverage](#113-coverage)
- [11.4 Benchmark Functions](#114-benchmark-functions)
- [11.5 Profiling](#115-profiling)
- [11.6 Example Functions](#116-example-functions)

---

## 11.1 The go test Tool

**File:** `11.1-the-go-test-tool.go`

**Learn:** `go test`, `-v`, `-count`, `-run`, table of packages

- [ ] A function `Add(a, b int) int` and `add_test.go` with `TestAdd`. Run `go test`, then `go test -v`.
- [ ] `go test -run TestAdd` vs `-run Nope` (no tests run).
- [ ] **Problem:** `go test ./...` from repo root. Fix whatever is not a buildable package yet (or skip empty folders).

---

## 11.2 Test Functions

**File:** `11.2-test-functions.go`

**Learn:** `func TestXxx(t *testing.T)`, `t.Error`/`Fatal`, table-driven tests, `t.Run`, helpers, `t.Helper`, golden files (light), `httptest` if useful

- [ ] Table-driven tests: slice of `{in, want}`. Loop with `t.Run(name, func(t *testing.T) { ... })`.
- [ ] `t.Fatalf` vs `t.Errorf` — one test with two bad cases: `Error` continues, `Fatal` stops. See `-v`.
- [ ] Compare slices/maps with a loop or `reflect.DeepEqual` (or `go-cmp` later; stdlib is enough).
- [ ] **Problem:** `word.Count(s string) map[string]int` with at least 5 table cases (empty, punctuation, unicode).
- [ ] **Problem:** `t.Helper()` in `assertEq`. Fail a test on purpose once to see the line number, then fix.
- [ ] **Problem:** extract `IntSet` or `tempconv` from earlier chapters into this folder (or test them in place) with table tests for the tricky methods.

---

## 11.3 Coverage

**File:** `11.3-coverage.go`

**Learn:** `go test -cover`, `-coverprofile`, `go tool cover -html`

- [ ] `go test -cover`.
- [ ] `go test -coverprofile=cover.out` then `go tool cover -func=cover.out`.
- [ ] `go tool cover -html=cover.out` (open the HTML). Find a red (uncovered) branch. Add a test. Re-run.
- [ ] **Problem:** get `word.Count` (or your target) to 100% or document the branch you refuse to test and why.

---

## 11.4 Benchmark Functions

**File:** `11.4-benchmark-functions.go`

**Learn:** `func BenchmarkXxx(b *testing.B)`, `for i := 0; i < b.N; i++`, `b.ResetTimer`, `-bench`, `-benchmem`

- [ ] `BenchmarkPopCount` vs a naive loop (bring `popcount` from Chapter 3 or rewrite).
- [ ] `go test -bench=. -benchmem`.
- [ ] **Problem:** benchmark `strings.Join` vs `+=` in a loop for 100 parts. Read allocations.
- [ ] **Problem:** `b.StopTimer` around setup if you allocate a big input once.

---

## 11.5 Profiling

**File:** `11.5-profiling.go`

**Learn:** `cpuprofile`, `memprofile`, `go tool pprof` (web or `top`)

- [ ] `go test -bench=. -cpuprofile=cpu.out` on a bench that actually burns CPU.
- [ ] `go tool pprof -top cpu.out` (or `pprof -http=:8080`). Note the hottest function.
- [ ] **Problem:** change the code to be faster (or slower on purpose), re-profile, write 5 lines in `notes.md` on what changed.

---

## 11.6 Example Functions

**File:** `11.6-example-functions.go`

**Learn:** `func ExampleXxx()`, output comment, `godoc` examples, compilable docs

- [ ] `func ExampleAdd()` that prints a result. Trailing comment `// Output: 3`.
- [ ] `go test` runs examples. Break the output — test fails.
- [ ] **Problem:** `ExampleCount` for your word counter showing the map print (sort keys so output is stable).
- [ ] **Problem:** run `go doc -all .` in that package and confirm the example appears.
