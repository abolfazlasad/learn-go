# Chapter 4 — Composite Types

**Goal:** arrays, slices, maps, structs, JSON, and templates — by building small data programs.

Write each section as its own file in this folder. Run one file at a time: `go run 4.1-arrays.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [4.1 Arrays](#41-arrays)
- [4.2 Slices](#42-slices)
- [4.3 Maps](#43-maps)
- [4.4 Structs](#44-structs)
- [4.5 JSON](#45-json)
- [4.6 Text and HTML Templates](#46-text-and-html-templates)

---

## 4.1 Arrays

**File:** `4.1-arrays.go`

**Learn:** `[N]T`, length is part of the type, value semantics, `[...]T`

- [ ] `var a [3]int`. Print `a`, `len(a)`, `a[0]`. Set `a[1]`.
- [ ] `b := [3]int{1, 2, 3}` and `c := [...]int{1, 2, 3}`. Compare `a == b`.
- [ ] Pass an array to a function, change an index inside, print before/after in `main`. Arrays copy.
- [ ] **Problem:** `[32]byte` SHA-256-looking hex printer: given 32 bytes, print 64 hex chars. Use a tiny hardcoded array.
- [ ] **Problem:** reverse an array in place via a pointer to the array, or by returning a new array. Show both.

---

## 4.2 Slices

**File:** `4.2-slices.go`

**Learn:** slice header (ptr, len, cap), `make`, `append`, slicing, shared backing array, `copy`

- [ ] From `s := []int{0,1,2,3,4,5}`, print slices `s[1:4]`, `s[:3]`, `s[3:]`, `s[:]`.
- [ ] Print `len` and `cap` after `make([]int, 0, 5)` and after several `append`s.
- [ ] Change `s[1:3][0] = 99` and show the original slice changed (shared array).
- [ ] `append` that exceeds cap — prove the original is *not* updated if you ignore the return value.
- [ ] `copy(dst, src)`.
- [ ] **Problem:** `reverse([]int)` in place. `rotate(s []int, n int)` left by n. `remove(s []int, i int) []int` without leaving a hole.
- [ ] **Problem:** `appendInt` written yourself (grow cap when needed). Compare with builtin `append`.
- [ ] **Problem:** nonempty filter — copy in place, return a shorter slice. Then a version that does not alias the original.

---

## 4.3 Maps

**File:** `4.3-maps.go`

**Learn:** `make(map[K]V)`, comma-ok, `delete`, ranging, zero value `nil` map, keys must be comparable

- [ ] Count word frequencies from args or stdin. Print the map.
- [ ] Look up a missing key. Then use `v, ok := m[k]`.
- [ ] `delete(m, k)`. Range prints in random order — run twice.
- [ ] `var m map[string]int` then `m["a"] = 1` — panic. Fix with `make`.
- [ ] **Problem:** `equal(a, b map[string]int) bool` (same keys and values).
- [ ] **Problem:** graph as `map[string][]string`. Function `addEdge`, `neighbors`, and `reachable(from, to)` with a visited set.
- [ ] **Problem:** sort printed keys (`[]string` + `sort.Strings`) so output is stable.

---

## 4.4 Structs

**File:** `4.4-structs.go`

**Learn:** struct literals, named vs positional, exported fields, embedding preview, pointers to structs

- [ ] `type Employee struct { ID int; Name string; Salary int }`. Build with named fields.
- [ ] `p := &Employee{...}`. Set `p.Salary` (pointer indirection is automatic).
- [ ] Compare two structs with `==` when all fields are comparable.
- [ ] Embed `Address` inside `Employee`. Set `e.City` via promotion. Also set `e.Address.City`.
- [ ] **Problem:** `type Point struct{ X, Y float64 }` with `distance(p, q Point) float64`.
- [ ] **Problem:** a `Tree` node (`Value int; Left, Right *Tree`). Insert and in-order print.

---

## 4.5 JSON

**File:** `4.5-json.go`

**Learn:** `encoding/json`, tags, `Marshal`/`Unmarshal`, anonymous structs, streaming decoder (light)

- [ ] Marshal a struct to JSON. Print the bytes as a string.
- [ ] Add `json:"name"` tags. Hide a field with `json:"-"`.
- [ ] Unmarshal a JSON string into a struct. Missing fields stay zero.
- [ ] Unmarshal into `map[string]any` or `[]any` and type-assert a value.
- [ ] **Problem:** movie list — marshal a slice of structs, write `movies.json`, read it back, print titles.
- [ ] **Problem:** given JSON with extra unknown fields, unmarshal without failing. Then a strict version if you want (`Decoder.DisallowUnknownFields`).

---

## 4.6 Text and HTML Templates

**File:** `4.6-text-and-html-templates.go`

**Learn:** `text/template`, `html/template`, `{{.Field}}`, `range`, `if`, actions vs escaping

- [ ] `text/template`: execute a template with a struct (`{{.Name}}` `{{.Age}}`).
- [ ] `{{range .Items}}...{{end}}` and `{{if .OK}}`.
- [ ] Same data with `html/template` and a field containing `<script>`. Compare output to `text/template`.
- [ ] **Problem:** render a report of word frequencies (from 4.3) as text, sorted by count.
- [ ] **Problem:** tiny HTML page listing those words. Open the file in a browser. Confirm tags in the data are escaped.
