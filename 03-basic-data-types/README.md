# Chapter 3 — Basic Data Types

**Goal:** integers, floats, complex, bools, strings, and constants — with programs that expose overflow, rounding, and encoding.

Write each section as its own file in this folder. Run one file at a time: `go run 3.1-integers.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [3.1 Integers](#31-integers)
- [3.2 Floating-Point Numbers](#32-floating-point-numbers)
- [3.3 Complex Numbers](#33-complex-numbers)
- [3.4 Booleans](#34-booleans)
- [3.5 Strings](#35-strings)
- [3.6 Constants](#36-constants)

---

## 3.1 Integers

**File:** `3.1-integers.go`

**Learn:** `int`, `int8`…`int64`, `uint`, `byte`, `rune`, overflow, bit operations, `fmt` verbs

- [ ] Print `math.MaxInt8`, `math.MaxUint8`, `math.MaxInt64`. Add 1 to a `uint8` at 255 and print (wrap).
- [ ] Show `int` vs `int64`. Convert explicitly. Try assigning one to the other without conversion.
- [ ] Bit ops: `& | ^ &^ << >>` on small numbers. Print in binary with `%08b`.
- [ ] **Problem:** `popcount(n uint64) int` — number of 1-bits. Test with 0, 1, 255, `math.MaxUint64`.
- [ ] **Problem:** convert a uint32 to dotted IPv4 (`0x7f000001` → `127.0.0.1`) using shifts and masks.

---

## 3.2 Floating-Point Numbers

**File:** `3.2-floating-point-numbers.go`

**Learn:** `float32`/`float64`, `NaN`, `Inf`, rounding, `math` package, formatting

- [ ] Print `1.0/0.0`, `-1.0/0.0`, `0.0/0.0`. Check with `math.IsInf` and `math.IsNaN`.
- [ ] Show `0.1+0.2 == 0.3` is false. Print with `%.20f`.
- [ ] Format the same float with `%f`, `%e`, `%g`.
- [ ] **Problem:** surface area and volume of a sphere from a radius arg. Use `math.Pi`.
- [ ] **Problem:** loop `x := 1.0; x <= 10; x += 0.1` vs looping on an int and converting. Count iterations. Explain the difference in a comment.

---

## 3.3 Complex Numbers

**File:** `3.3-complex-numbers.go`

**Learn:** `complex64`/`complex128`, `real`, `imag`, `complex`, arithmetic

- [ ] Build `z := 1+2i`. Print `real(z)`, `imag(z)`, `z * z`.
- [ ] Use `complex(a, b)` from two floats.
- [ ] **Problem:** mandelbrot iteration count for one point `c` (z = z²+c until |z|>2 or max iter). Print the count. Optional: emit ASCII art for a tiny grid.

---

## 3.4 Booleans

**File:** `3.4-booleans.go`

**Learn:** `bool`, `&& || !`, short-circuit, no implicit conversion from int

- [ ] Try `var b bool = 1` — must fail. Use comparisons instead.
- [ ] Write `true && f()` and `false && f()` where `f` prints. Prove short-circuit.
- [ ] Same for `||`.
- [ ] **Problem:** `isLeap(year int) bool` using the Gregorian rule. Check 1900, 2000, 2024, 2025.

---

## 3.5 Strings

**File:** `3.5-strings.go`

**Learn:** immutable strings, bytes vs runes, `range` over string, `utf8`, slicing, `strings` package

- [ ] `s := "hello"`. Print `len(s)`, `s[0]`, `s[1:4]`. Try `s[0] = 'H'` — must not compile.
- [ ] `s := "سلام"`. Print `len(s)` vs `utf8.RuneCountInString(s)`. Range with `for i, r := range s`.
- [ ] Convert `[]byte(s)` and `string(bytes)`. Convert `[]rune(s)`.
- [ ] Use `strings.Contains`, `HasPrefix`, `Join`, `Split`, `ToUpper`.
- [ ] **Problem:** `basename(path string) string` — strip directory and optional suffix (like the book idea, written from scratch).
- [ ] **Problem:** count runes, bytes, and words on stdin. Words = `strings.Fields`.
- [ ] **Problem:** comma-separate an integer string (`12345` → `12,345`) without `fmt` commas.

---

## 3.6 Constants

**File:** `3.6-constants.go`

**Learn:** `const`, `iota`, untyped constants, typed vs untyped

- [ ] `const` of int, float, string. Try `const x = time.Now()` — must fail.
- [ ] `const ( a = iota; b; c )`. Print 0,1,2. Then a bit-flag set with `1 << iota`.
- [ ] Show an untyped numeric constant assigned to `int` and to `float64`.
- [ ] **Problem:** define `KB, MB, GB, TB` as `iota`-based powers of 1024. Print them.
- [ ] **Problem:** typed const `type Day int` with `Sunday = iota` … `Saturday`. Function `weekend(d Day) bool`.
