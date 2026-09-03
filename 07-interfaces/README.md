# Chapter 7 — Interfaces

**Goal:** contracts, satisfaction, interface values, standard interfaces, type assertions, type switches.

Do not skip 7.1–7.3. Later sections are where interfaces click.

Write each section as its own file in this folder. Run one file at a time: `go run 7.1-interfaces-as-contracts.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [7.1 Interfaces as Contracts](#71-interfaces-as-contracts)
- [7.2 Interface Types](#72-interface-types)
- [7.3 Interface Satisfaction](#73-interface-satisfaction)
- [7.4 Parsing Flags with flag.Value](#74-parsing-flags-with-flagvalue)
- [7.5 Interface Values](#75-interface-values)
- [7.6 Sorting with sort.Interface](#76-sorting-with-sortinterface)
- [7.7 The http.Handler Interface](#77-the-httphandler-interface)
- [7.8 The error Interface](#78-the-error-interface)
- [7.9 Example: Expression Evaluator](#79-example-expression-evaluator)
- [7.10 Type Assertions](#710-type-assertions)
- [7.11 Discriminating Errors with Type Assertions](#711-discriminating-errors-with-type-assertions)
- [7.12 Querying Behaviors with Interface Type Assertions](#712-querying-behaviors-with-interface-type-assertions)
- [7.13 Type Switches](#713-type-switches)
- [7.14 Example: Token-Based XML Decoding](#714-example-token-based-xml-decoding)
- [7.15 A Few Words of Advice](#715-a-few-words-of-advice)

---

## 7.1 Interfaces as Contracts

**File:** `7.1-interfaces-as-contracts.go`

**Learn:** a method set is the contract; callers depend on behavior, not concrete type

- [ ] `type Writer interface { Write([]byte) (int, error) }`. Implement it on `ByteCounter int` that counts bytes.
- [ ] Pass `*ByteCounter` to `fmt.Fprintf`. Print the count.
- [ ] **Problem:** `type Reader interface { Read([]byte) (int, error) }`. Implement a `LimitReader` that stops after N bytes from an inner `io.Reader`.
- [ ] **Problem:** write a function `dump(w io.Writer, s string)` that only uses the interface. Call it with `os.Stdout` and with `bytes.Buffer`.

---

## 7.2 Interface Types

**File:** `7.2-interface-types.go`

**Learn:** embedding interfaces, empty interface `any`, standard library interfaces

- [ ] `type ReadWriter interface { io.Reader; io.Writer }`. Show `*bytes.Buffer` works.
- [ ] Function `func printAny(x any)` using `%T` and `%v`.
- [ ] **Problem:** define `Stringer` yourself (`String() string`) and a `type IP [4]byte` that implements it. Print via `fmt.Println`.

---

## 7.3 Interface Satisfaction

**File:** `7.3-interface-satisfaction.go`

**Learn:** implicit implementation, pointer vs value method sets, compile-time check

- [ ] `var _ io.Writer = (*ByteCounter)(nil)` compile-time assertion.
- [ ] Type with a pointer-receiver `Write`. Assign `T{}` to `io.Writer` — should fail. Assign `&T{}` — should work.
- [ ] **Problem:** two types, `UpperWriter` and `CountWriter`, both `io.Writer`. A function `tee(w io.Writer, s string)` writes to either without knowing which.

---

## 7.4 Parsing Flags with flag.Value

**File:** `7.4-parsing-flags-with-flag-value.go`

**Learn:** `flag.Value` (`String`, `Set`), custom flag types

- [ ] Custom `type Celsius float64` implementing `flag.Value`. `flag.Var(&c, "temp", ...)`.
- [ ] Parse `-temp 100C` and `-temp 212F` (or similar). Print Celsius.
- [ ] **Problem:** `type Strings []string` flag that can be repeated (`-item a -item b`) collecting a slice.

---

## 7.5 Interface Values

**File:** `7.5-interface-values.go`

**Learn:** (dynamic type, dynamic value), nil interface vs interface holding nil pointer, `%T`

- [ ] `var w io.Writer`. Print `w == nil`. Assign `os.Stdout`. Print `%T`.
- [ ] `var b *bytes.Buffer; w = b`. Then `w.Write(...)` — panic. Fix by assigning a non-nil `*bytes.Buffer`.
- [ ] **Problem:** function `isNil(w io.Writer) bool` that reports both “interface is nil” and “dynamic value is nil” (use `w == nil` plus reflect later, or compare documented behavior in comments and a type assert).
- [ ] **Problem:** table of assignments (`nil`, `*os.File(nil)`, `os.Stdout`) and whether `Write` panics.

---

## 7.6 Sorting with sort.Interface

**File:** `7.6-sorting-with-sort-interface.go`

**Learn:** `Len`, `Less`, `Swap`; `sort.Sort`; `sort.Slice` as the modern shortcut

- [ ] `type StringSlice []string` implementing `sort.Interface`. Sort names.
- [ ] Sort a slice of structs by a field using a named type `BySalary []Employee`.
- [ ] Same sort with `sort.Slice`.
- [ ] **Problem:** sort tracks/songs by artist, then by year, using a `customSort` struct that holds the slice and a `less` function (or two named types).
- [ ] **Problem:** reverse sort with `sort.Reverse`.

---

## 7.7 The http.Handler Interface

**File:** `7.7-the-http-handler-interface.go`

**Learn:** `ServeHTTP`, `http.HandlerFunc` adapter, `http.ServeMux`

- [ ] Type `type counter int` with `ServeHTTP`. Register it. Hit `/` and see the count change.
- [ ] Convert a function with `http.HandlerFunc(fn)`.
- [ ] `http.NewServeMux()`, handle `/hello` and `/time`.
- [ ] **Problem:** `type database map[string]int` as a handler: `GET /list` lists items, `GET /price?item=x` returns a price.
- [ ] **Problem:** wrap a handler with a logger middleware: type `logHandler struct{ inner http.Handler }` that logs method + path then calls `inner`.

---

## 7.8 The error Interface

**File:** `7.8-the-error-interface.go`

**Learn:** `Error() string`, sentinel errors, custom error types with extra fields

- [ ] `var errNotFound = errors.New("not found")`. Return it. Compare with `==`.
- [ ] `type PathError struct { Op, Path string; Err error }` with `Error()`.
- [ ] **Problem:** `type StatusError struct { Code int; Msg string }`. Function that returns it. Caller prints `Code`.
- [ ] **Problem:** wrap with `fmt.Errorf("load config: %w", err)` and unwrap with `errors.Is` / `errors.As`.

---

## 7.9 Example: Expression Evaluator

**File:** `7.9-example-expression-evaluator.go`

**Learn:** an interface with several concrete types (`Var`, `Literal`, `Unary`, `Binary`)

- [ ] `type Expr interface { Eval(env map[string]float64) float64 }`.
- [ ] Implement `Var`, `Literal`, `Unary` (`+` `-`), `Binary` (`+ - * /`).
- [ ] **Problem:** evaluate `a + 2 * b` with `env{"a": 3, "b": 4}` → 11 (or the parse tree you build by hand).
- [ ] **Problem:** `Check() error` on `Expr` that rejects unknown variables and division by zero constants if you want extra credit.
- [ ] Optional later: parse from a string. Hand-built trees are enough here.

---

## 7.10 Type Assertions

**File:** `7.10-type-assertions.go`

**Learn:** `x.(T)`, comma-ok, panic if wrong and no `ok`

- [ ] `var i any = "hello"; s := i.(string)`. Then `i.(int)` with `ok`. Then without `ok` and see panic.
- [ ] **Problem:** `func stringify(v any) string` using assertions for `string`, `int`, `fmt.Stringer`, default `%v`.

---

## 7.11 Discriminating Errors with Type Assertions

**File:** `7.11-discriminating-errors-with-type-assertions.go`

**Learn:** `err.(*os.PathError)`, `errors.As`, handling by type

- [ ] Open a missing file. Switch on the error type (`*os.PathError` vs others).
- [ ] Same with `errors.As`.
- [ ] **Problem:** function `explain(err error) string` that returns `not found`, `permission`, or `other` based on `os.IsNotExist`, `os.IsPermission`, or type.

---

## 7.12 Querying Behaviors with Interface Type Assertions

**File:** `7.12-querying-behaviors-with-interface-type-assertions.go`

**Learn:** assert an optional interface (`io.WriterTo`, `http.Flusher`, custom `Close() error`)

- [ ] `type Closer interface { Close() error }`. If `w` also implements `Closer`, close it; else skip.
- [ ] **Problem:** `writeAll(w io.Writer, p []byte)` — if `w` has `WriteString(string) (int, error)`, use that for a string conversion path; else `Write`.
- [ ] **Problem:** HTTP handler that type-asserts `http.Flusher` and flushes after each line of a slow response.

---

## 7.13 Type Switches

**File:** `7.13-type-switches.go`

**Learn:** `switch v := x.(type)`

- [ ] Type switch on `any` for `int`, `string`, `[]byte`, default.
- [ ] **Problem:** `sqlLiteral(v any) string` that quotes strings, prints numbers, `NULL` for nil, error for unknown types.
- [ ] **Problem:** rewrite `stringify` from 7.10 as a type switch.

---

## 7.14 Example: Token-Based XML Decoding

**File:** `7.14-example-token-based-xml-decoding.go`

**Learn:** `encoding/xml` tokens, type switch on `xml.StartElement`, `CharData`, `EndElement`

- [ ] Decode a small XML string token by token. Print each token's type.
- [ ] Collect text of a given element name (`<item>...</item>`).
- [ ] **Problem:** XML → nested outline printed with indentation (stack of element names).
- [ ] **Problem:** extract all `href` attributes from `<a>` tags in an HTML/XML snippet.

---

## 7.15 A Few Words of Advice

**File:** `7.15-a-few-words-of-advice.go`

**Learn:** keep interfaces small; accept interfaces, return concrete types

- [ ] Refactor an earlier function that took `*bytes.Buffer` so it takes `io.Writer` instead. Call it with `os.Stdout` and a file.
- [ ] **Problem:** find one place in your Chapter 7 code that used `any` lazily. Replace it with a small interface or a concrete type. Write 5 lines in `notes.md` on why.
- [ ] **Problem:** define a one-method interface you actually need (`Namer`, `Validator`, …) and two types that satisfy it. Do not add methods “for the future.”
