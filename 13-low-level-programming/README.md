# Chapter 13 — Low-Level Programming

**Goal:** `unsafe` and `cgo` exist. Use them when you must, with tests and comments. Most application code never needs this chapter’s tools.

If a todo feels dangerous, that is the lesson. Prefer the safe API unless the exercise requires otherwise.

Write each section as its own file in this folder. Run one file at a time: `go run 13.1-unsafe-sizeof-alignof-and-offsetof.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [13.1 unsafe.Sizeof, Alignof, and Offsetof](#131-unsafesizeof-alignof-and-offsetof)
- [13.2 unsafe.Pointer](#132-unsafepointer)
- [13.3 Example: Deep Equivalence](#133-example-deep-equivalence)
- [13.4 Calling C Code with cgo](#134-calling-c-code-with-cgo)
- [13.5 Another Word of Caution](#135-another-word-of-caution)

---

## 13.1 unsafe.Sizeof, Alignof, and Offsetof

**File:** `13.1-unsafe-sizeof-alignof-and-offsetof.go`

**Learn:** size and alignment of types; field offsets; padding

- [ ] Print `unsafe.Sizeof` for `int`, `int32`, `int64`, `string`, `[]int`, a small struct.
- [ ] Two structs with the same fields in different order. Compare `Sizeof`. Reorder to shrink padding.
- [ ] `unsafe.Offsetof(s.Field)` for each field. Match against a mental layout.
- [ ] **Problem:** `type S struct { A bool; B float64; C int32 }`. Draw the layout in a comment (offsets + padding). Then reorder fields and print the new size.
- [ ] **Problem:** `Alignof` of the struct vs of each field.

---

## 13.2 unsafe.Pointer

**File:** `13.2-unsafe-pointer.go`

**Learn:** `unsafe.Pointer` conversions, the allowed patterns (via `uintptr` briefly), aliasing risks

- [ ] Convert `*T` → `unsafe.Pointer` → `*U` for two types of the same size (e.g. `float64` bits as `uint64`). Print bits. Prefer `math.Float64bits` after you see why.
- [ ] **Problem:** read the `Float64bits` source (`go doc` / stdlib). Reimplement a one-liner with `unsafe`, then delete it and use `math`.
- [ ] **Problem:** `notes.md` — which conversions are allowed; why `uintptr` held across a call is a GC bug. No need to write the bug.

---

## 13.3 Example: Deep Equivalence

**File:** `13.3-example-deep-equivalence.go`

**Learn:** comparing values reflection cannot (`unsafe` for unexported / cycles) — or a careful `DeepEqual` of your own

- [ ] Write `Equal(x, y any) bool` using reflect (like `reflect.DeepEqual`). Handle numbers, strings, slices, maps, pointers, structs.
- [ ] Cycle: two linked lists that point at themselves. Must not infinite-loop (seen map of addresses).
- [ ] **Problem:** compare two structs with a slice field. Then NaN floats — decide and test the behavior.
- [ ] **Problem:** tests vs `reflect.DeepEqual` for 8 cases. Document any difference.

---

## 13.4 Calling C Code with cgo

**File:** `13.4-calling-c-code-with-cgo.go`

**Learn:** `import "C"`, `// #include`, calling a C function, `C.CString` / `C.free`

Skip this section on a machine without a C compiler if `cgo` is disabled (`CGO_ENABLED=0`). Check: `go env CGO_ENABLED`.

- [ ] `// #include <stdio.h>` and `import "C"`. Call `C.puts`.
- [ ] Pass a Go string: `cs := C.CString(s); defer C.free(unsafe.Pointer(cs))`.
- [ ] **Problem:** C function `int add(int a, int b)` in a comment block (or `add.c`). Call from Go. Test it.
- [ ] **Problem:** `notes.md` — costs of cgo (and when to use a pure Go library instead).

---

## 13.5 Another Word of Caution

**File:** `13.5-another-word-of-caution.go`

**Learn:** unsafe/cgo opt out of Go’s guarantees

- [ ] Re-read your Chapter 13 programs. Add a 3-line comment above every `unsafe` use: what invariant makes it correct.
- [ ] **Problem:** `notes.md` — “I will use unsafe/cgo when …” (empty is a valid answer if you cannot name a real case).
- [ ] **Problem:** pick one earlier chapter type (`IntSet`, JSON, templates). Implement a feature *without* unsafe. That is the default path.
