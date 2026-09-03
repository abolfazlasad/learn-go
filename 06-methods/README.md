# Chapter 6 — Methods

**Goal:** methods, pointer vs value receivers, embedding, method values, a real type, encapsulation.

Write each section as its own file in this folder. Run one file at a time: `go run 6.1-method-declarations.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [6.1 Method Declarations](#61-method-declarations)
- [6.2 Methods with a Pointer Receiver](#62-methods-with-a-pointer-receiver)
- [6.3 Composing Types by Struct Embedding](#63-composing-types-by-struct-embedding)
- [6.4 Method Values and Expressions](#64-method-values-and-expressions)
- [6.5 Example: Bit Vector Type](#65-example-bit-vector-type)
- [6.6 Encapsulation](#66-encapsulation)

---

## 6.1 Method Declarations

**File:** `6.1-method-declarations.go`

**Learn:** `func (t T) Method(...)`, receiver, methods vs functions

- [ ] `type Point struct{ X, Y float64 }` with `func (p Point) Distance(q Point) float64`.
- [ ] Call `p.Distance(q)`. Also call the method with a pointer variable — Go takes the address if needed for the other kind later.
- [ ] **Problem:** `type IntList []int` with `Sum() int` and `Max() (int, bool)`.
- [ ] **Problem:** same `Distance` as a package function `Distance(p, q Point)`. Compare call sites.

---

## 6.2 Methods with a Pointer Receiver

**File:** `6.2-methods-with-a-pointer-receiver.go`

**Learn:** `(p *T)` when you mutate or avoid copies; nil receivers; mixing receiver types

- [ ] `func (p *Point) ScaleBy(factor float64)` that mutates. Show a value receiver would not persist.
- [ ] Call `ScaleBy` on a `Point` value (addressable). Try on a temporary `Point{1,2}.ScaleBy(2)` — note the compile error if the receiver is a pointer.
- [ ] Method on a nil receiver: `func (l *IntList) Sum() int` that returns 0 if `l == nil`.
- [ ] **Problem:** `type Counter struct{ n int }` with `Inc()`, `Value() int`. `Inc` must use a pointer. Prove `Value` can use either, and pick one consistently.
- [ ] **Problem:** document in comments: if any method uses a pointer receiver, all mutating ones should.

---

## 6.3 Composing Types by Struct Embedding

**File:** `6.3-composing-types-by-struct-embedding.go`

**Learn:** anonymous fields, promoted methods, wrapping vs inheritance

- [ ] `type ColoredPoint struct { Point; Color string }`. Call `cp.Distance` (promoted).
- [ ] Embed `*Point` instead. Nil inner pointer + method call — see the panic, then nil-check.
- [ ] **Problem:** `type Logger struct{ Prefix string }` with `Log(msg string)`. Embed it in `Server`. `Server` should log with a prefix without reimplementing `Log`.
- [ ] **Problem:** override a promoted method by defining the same name on the outer type. Call both `s.Log` and `s.Logger.Log`.

---

## 6.4 Method Values and Expressions

**File:** `6.4-method-values-and-expressions.go`

**Learn:** `p.Distance` as a value, `Point.Distance` as a function

- [ ] `d := p.Distance` then `d(q)`. The receiver is bound.
- [ ] `f := Point.Distance` then `f(p, q)`.
- [ ] **Problem:** a slice of `func()` built from method values (`p.ScaleBy` bound to different points) and run them all.
- [ ] **Problem:** pass `p.Distance` into `apply` from Chapter 5 if you still have that idea, or into `sort.Slice` via a method value on a wrapper.

---

## 6.5 Example: Bit Vector Type

**File:** `6.5-example-bit-vector-type.go`

**Learn:** a small real type with methods; `[]uint64` as bits

- [ ] `type IntSet struct { words []uint64 }` with `Has(x int) bool`, `Add(x int)`, `String() string`.
- [ ] `Add` must grow `words` as needed. `Has` on a value past the end is false, not panic.
- [ ] **Problem:** `UnionWith(t *IntSet)`, `Len()`, `Remove(x int)`, `Clear()`, `Copy() *IntSet`.
- [ ] **Problem:** `Elems() []int` returning all elements in order.
- [ ] Write a `main` that adds 1, 144, 9, 42 and prints the set.

---

## 6.6 Encapsulation

**File:** `6.6-encapsulation.go`

**Learn:** exported identifiers, unexported fields, constructors, another package as the client

- [ ] This section needs two packages, so a small `6.6-counter/` directory is allowed: unexported field `n`, plus `6.6-main.go` that imports it.
- [ ] Client cannot set `n` directly. Provide `New()`, `Inc()`, `Value()`.
- [ ] **Problem:** `geometry` package: `Point` with unexported fields, `NewPoint(x,y)`, `X()`, `Y()`, `Move`. Main in another package uses only the API.
- [ ] **Problem:** try to access an unexported field from `main` — paste the compiler error in a comment, then delete the illegal line.
