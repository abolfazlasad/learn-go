# Chapter 12 — Reflection

**Goal:** `reflect` when you cannot know types at compile time (fmt, json, your own generic-looking helpers). Then stop using it when a type switch or generics would do.

Go 1.18+ has generics; still do this chapter — `encoding/json` is reflect underneath, and the book’s tools are still how you debug surprising APIs.

Write each section as its own file in this folder. Run one file at a time: `go run 12.1-why-reflection.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [12.1 Why Reflection?](#121-why-reflection)
- [12.2 reflect.Type and reflect.Value](#122-reflecttype-and-reflectvalue)
- [12.3 Display, a Recursive Value Printer](#123-display-a-recursive-value-printer)
- [12.4 Example: Encoding S-Expressions](#124-example-encoding-s-expressions)
- [12.5 Setting Variables with reflect.Value](#125-setting-variables-with-reflectvalue)
- [12.6 Example: Decoding S-Expressions](#126-example-decoding-s-expressions)
- [12.7 Accessing Struct Field Tags](#127-accessing-struct-field-tags)
- [12.8 Displaying the Methods of a Type](#128-displaying-the-methods-of-a-type)
- [12.9 A Word of Caution](#129-a-word-of-caution)

---

## 12.1 Why Reflection?

**File:** `12.1-why-reflection.go`

**Learn:** `fmt` and `json` need to inspect unknown values; the cost is safety and speed

- [ ] Write `printAny(v any)` with a type switch covering `int`, `string`, default. Then try a `struct` — you need reflect or `%v`.
- [ ] **Problem:** `notes.md` — three cases where reflect is justified, three where it is not.

---

## 12.2 reflect.Type and reflect.Value

**File:** `12.2-reflect-type-and-reflect-value.go`

**Learn:** `reflect.TypeOf`, `ValueOf`, `Kind`, `Int()`, `String()`, `Elem`, `Interface()`

- [ ] Print `TypeOf` and `Kind` for `int`, `*int`, `[]int`, a struct, an interface holding an int.
- [ ] `ValueOf(3).Int()` vs `ValueOf("hi").String()`. Wrong kind → panic; recover or choose correctly.
- [ ] `ValueOf(&x).Elem()` to get the settable `x`.
- [ ] **Problem:** `describe(v any)` prints type, kind, and value for at least 6 different kinds.

---

## 12.3 Display, a Recursive Value Printer

**File:** `12.3-display-a-recursive-value-printer.go`

**Learn:** walk structs, slices, maps, pointers, cycles (optional), unexported fields

- [ ] Recursive `Display(name string, x any)` that prints a tree for structs and slices.
- [ ] Pointers: print type and pointed-to value. Nil pointer is obvious.
- [ ] **Problem:** pretty-print a nested struct (`Movie` with `[]Actor`).
- [ ] **Problem:** a slice that contains a pointer back to itself or a map cycle — either detect cycles or document that you skip them.

---

## 12.4 Example: Encoding S-Expressions

**File:** `12.4-example-encoding-s-expressions.go`

**Learn:** marshal arbitrary values to a text format with reflect

- [ ] `Marshal(v any) ([]byte, error)` for ints, strings, structs → `(Movie (Title "x") (Year 1978))` style, or a simpler Lisp-like form you define.
- [ ] Slices as lists. Skip unexported fields.
- [ ] **Problem:** marshal the `Movie` from 12.3. Compare by eye to JSON.
- [ ] **Problem:** `bool` and `[]string` support.

---

## 12.5 Setting Variables with reflect.Value

**File:** `12.5-setting-variables-with-reflect-value.go`

**Learn:** `CanSet`, `SetInt`, `SetString`, addressability

- [ ] `reflect.ValueOf(x).CanSet()` is false. `ValueOf(&x).Elem().CanSet()` is true. Set a new int.
- [ ] Try `Set` on a non-addressable value — catch the panic, then fix.
- [ ] **Problem:** `func setField(ptr any, name string, val any) error` that sets an exported struct field by name.
- [ ] **Problem:** refuse unexported fields with a clear error.

---

## 12.6 Example: Decoding S-Expressions

**File:** `12.6-example-decoding-s-expressions.go`

**Learn:** unmarshal into a pointer using reflect; tokenizer

- [ ] Tokenize a tiny S-expression (`Scanner` or hand-rolled).
- [ ] `Unmarshal(data []byte, dst any)` into a struct pointer.
- [ ] **Problem:** round-trip: marshal then unmarshal `Movie`. `reflect.DeepEqual` the result.
- [ ] **Problem:** good errors on malformed input (not panic).

---

## 12.7 Accessing Struct Field Tags

**File:** `12.7-accessing-struct-field-tags.go`

**Learn:** `StructTag.Get("json")`, `http`, `db` tags

- [ ] Walk a struct type, print each field name, type, and `json` tag.
- [ ] **Problem:** mini-validator: tag `validate:"required"` and a `Validate(v any) error` that checks zero values on tagged fields.
- [ ] **Problem:** custom marshal that uses `sexpr:"name"` tags for the key instead of the field name.

---

## 12.8 Displaying the Methods of a Type

**File:** `12.8-displaying-the-methods-of-a-type.go`

**Learn:** `Type.NumMethod`, `Method(i)`, `PkgPath` for unexported

- [ ] Print all methods of `os.File` or `bytes.Buffer` (name + signature).
- [ ] Print methods of your `Point` type from Chapter 6.
- [ ] **Problem:** `printMethods(v any)` used like `fmt` debugging. Include pointer type vs value type method sets (`T` vs `*T`).

---

## 12.9 A Word of Caution

**File:** `12.9-a-word-of-caution.go`

**Learn:** reflect hides bugs until runtime; prefer interfaces and generics for new code

- [ ] Rewrite `describe` from 12.2 as a generic function `func describe[T any](v T)` that only prints `%T %v` — notice what you *lost*.
- [ ] **Problem:** pick one helper from this chapter that should stay reflect (`json`-like marshal) and one that should not (`setField` for a known struct). `notes.md` with the choice.
- [ ] **Problem:** add a test that would have been a compile error without reflect (wrong field name string). See it fail at runtime. That is the caution.
