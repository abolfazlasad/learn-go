# Chapter 2 — Program Structure

**Goal:** names, declarations, variables, assignment, types, packages, and scope — by writing code that would not compile if you get them wrong.

Write each section as its own file in this folder. Run one file at a time: `go run 2.1-names.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [2.1 Names](#21-names)
- [2.2 Declarations](#22-declarations)
- [2.3 Variables](#23-variables)
- [2.4 Assignments](#24-assignments)
- [2.5 Type Declarations](#25-type-declarations)
- [2.6 Packages and Files](#26-packages-and-files)
- [2.7 Scope](#27-scope)

---

## 2.1 Names

**File:** `2.1-names.go`

**Learn:** exported vs unexported names, letters/digits/underscore, idiomatic mixedCaps

- [ ] In one file, declare `var maxRetry int` and `var MaxRetry int`. Print both. Which one is exported?
- [ ] Try `var 1x int` and `var type int`. Read the compiler errors.
- [ ] Rename a function from `compute_sum` to `computeSum`. Call it.
- [ ] **Problem:** write a package-level comment and a function comment in `godoc` style. Run `go doc` on that function.

---

## 2.2 Declarations

**File:** `2.2-declarations.go`

**Learn:** `var`, `const`, `type`, `func` at package level; declaration order

- [ ] Declare a package-level `const`, a `var`, a `type`, and a `func`. Use all four in `main`.
- [ ] Call a function that is declared *below* `main`. Confirm that works.
- [ ] **Problem:** a file with at least two consts, two vars, one named type, and two functions. The program should print a short report using all of them.

---

## 2.3 Variables

**File:** `2.3-variables.go`

**Learn:** `var x T`, `var x T = v`, `var x = v`, `x := v`, zero values, pointers, `new`

- [ ] For `int`, `string`, `bool`, print the value of an uninitialized `var`.
- [ ] Show three ways to declare an `int` set to 10. Print types with `%T`.
- [ ] Use `:=` inside `main`. Try `:=` at package level — it must fail.
- [ ] Take the address of a variable with `&`. Dereference with `*`. Change the value through the pointer.
- [ ] Use `new(int)`, set `*p = 42`, print.
- [ ] **Problem:** `swap(a, b *int)` that swaps through pointers. Prove it in `main`.
- [ ] **Problem:** a function that returns a pointer to a local variable. Print it in `main`. Does it work? Write a comment explaining why.

---

## 2.4 Assignments

**File:** `2.4-assignments.go`

**Learn:** assignment, tuple assignment, `++` / `--` as statements, assignment operators

- [ ] Swap two ints with a temp, then with `a, b = b, a`.
- [ ] Use `+=`, `-=`, `*=` in a short loop.
- [ ] Show that `i++` is a statement, not an expression. `x = i++` must not compile.
- [ ] **Problem:** given three ints `x, y, z`, rotate left so `x,y,z` become `y,z,x` in one statement.
- [ ] **Problem:** parse two integers from args and print quotient and remainder. Division by zero → error message, no panic.

---

## 2.5 Type Declarations

**File:** `2.5-type-declarations.go`

**Learn:** `type T underlying`, conversion `T(v)`, named types are distinct

- [ ] `type Celsius float64` and `type Fahrenheit float64`. Convert between them with a function, not by assigning directly.
- [ ] Try `var c Celsius = 100; var f Fahrenheit = c` — it must not compile. Fix with an explicit conversion *after* applying the formula.
- [ ] **Problem:** `CToF` and `FToC`. Print freezing and boiling in both units.
- [ ] **Problem:** `type Percentage int`. Reject values outside 0–100 in a function that returns `(Percentage, error)`.

---

## 2.6 Packages and Files

**File:** `2.6-packages-and-files.go`

**Learn:** one package, many files; import paths; initializer `init` (lightly)

- [ ] Split a tiny program across two files with the same prefix (`2.6-temp.go` + `2.6-tempconv.go`). Run both: `go run 2.6-temp.go 2.6-tempconv.go`.
- [ ] Put `package main` in one file and a different package name in another — read the error.
- [ ] Import `fmt` and `os`. Call something from each.
- [ ] **Problem:** a two-file program: `2.6-main.go` parses a number from args; `2.6-convert.go` holds `CToF`/`FToC`. Main prints both conversions.

---

## 2.7 Scope

**File:** `2.7-scope.go`

**Learn:** block scope, inner names shadow outer names, `if`/`for` scopes, package vs file

- [ ] Declare `x := 1` in `main`, then an inner block `{ x := 2; ... }`. Print `x` inside and after the block.
- [ ] `if x := os.Args; len(x) > 1 { ... }` then try to use `x` after the `if`. It must not compile.
- [ ] Two files, both `package main`, each with `var n = ...`. See what happens. Then use different names.
- [ ] **Problem:** a function with three nested scopes each defining `n`. Print `n` at each level so the shadowing is obvious.
- [ ] **Problem:** fix a program (write it broken first) where a `err :=` inside an `if` hides an outer `err` you still need.
