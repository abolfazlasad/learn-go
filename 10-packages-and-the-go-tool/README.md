# Chapter 10 — Packages and the Go Tool

**Goal:** import paths, package names, blank imports, and the `go` command by using them, not by memorizing flags.

This chapter is tool-heavy. Every todo is a command you run plus a small package you write.

Write each section as its own file in this folder. Run one file at a time: `go run 10.1-introduction.go` (several `package main` files cannot be built together with `go run .`).

When a section needs more than one file, keep the same prefix (`2.6-convert.go`, `2.6-main.go`) or a small sub-package only if the task says so.

## Sections

- [10.1 Introduction](#101-introduction)
- [10.2 Import Paths](#102-import-paths)
- [10.3 The Package Declaration](#103-the-package-declaration)
- [10.4 Import Declarations](#104-import-declarations)
- [10.5 Blank Imports](#105-blank-imports)
- [10.6 Packages and Naming](#106-packages-and-naming)
- [10.7 The Go Tool](#107-the-go-tool)

---

## 10.1 Introduction

**File:** `10.1-introduction.go`

**Learn:** a package is a directory of `.go` files with the same `package` name

- [ ] `go mod init` at the **repo root** if you have not yet. Confirm `go.mod` exists.
- [ ] Create a tiny library package in this folder (or `hello/` subfolder) and a `main` that imports it via the module path.
- [ ] **Problem:** `greet.Hello(name string) string` in package `greet`. Main prints `greet.Hello("abolfazl")`.

---

## 10.2 Import Paths

**File:** `10.2-import-paths.go`

**Learn:** module path + directory, stdlib paths, no relative imports in modern modules

- [ ] Import `fmt` and your own package. Print both import paths in comments.
- [ ] Try `import "./greet"` — see why it fails (or is discouraged). Use the module path instead.
- [ ] **Problem:** two packages, `tempconv` and `main`, with a correct import path. `go build ./...` from repo root succeeds.

---

## 10.3 The Package Declaration

**File:** `10.3-the-package-declaration.go`

**Learn:** `package foo` vs directory name, `package main` is special

- [ ] Directory `word/` with `package word`. Import it as `.../word`.
- [ ] Two files in the same directory with different `package` names — `go build` error.
- [ ] **Problem:** a non-main package that you run tests on (`go test`) vs a `main` you `go run`.

---

## 10.4 Import Declarations

**File:** `10.4-import-declarations.go`

**Learn:** grouped imports, aliases, `.` import (avoid), unused import is an error

- [ ] Import `fmt` and `strings`. Leave one unused — fix the compiler error.
- [ ] Alias: `import f "fmt"` then `f.Println`.
- [ ] **Problem:** write a file with a standard import block: stdlib, then blank line, then your module. Run `gofmt` / `goimports` if you have it.

---

## 10.5 Blank Imports

**File:** `10.5-blank-imports.go`

**Learn:** `import _ "image/png"` for side effects (`init`)

- [ ] Decode an image. Without blank-importing the format, PNG/JPEG decode fails. Add `_ "image/png"` (and jpeg if needed).
- [ ] Write a package with `func init() { fmt.Println("init") }` and blank-import it from `main`. See `init` run.
- [ ] **Problem:** `image.Decode` a PNG you generate or check in. Print width/height. Remove the blank import and record the error.

---

## 10.6 Packages and Naming

**File:** `10.6-packages-and-naming.go`

**Learn:** package name is the default identifier; avoid stutter; good exported names

- [ ] Bad: `tempconv.TempConvCToF`. Good: `tempconv.CToF`. Refactor a small package to the good style.
- [ ] **Problem:** package `http` is already taken by stdlib. Name yours `client` or `fetch` and export `Get`. Main calls `fetch.Get`.
- [ ] **Problem:** `notes.md` — 5 naming rules you will follow (from this section + what `gofmt` already forced).

---

## 10.7 The Go Tool

**File:** `10.7-the-go-tool.go`

**Learn:** `go build`, `run`, `test`, `fmt`, `list`, `doc`, `vet`, `mod`, `env`

- [ ] From repo root: `go fmt ./...`, `go vet ./...`, `go list ./...`.
- [ ] `go env GOPATH GOROOT GOOS GOARCH`.
- [ ] `go doc fmt.Fprintf` and `go doc -src fmt.Fprintf` (skim).
- [ ] `go build -o /tmp/hello 01-tutorial/1.1-hello-world.go` (after Chapter 1 exists).
- [ ] **Problem:** write a one-screen cheat sheet `GOTOL.md` in this folder with the 10 commands you actually used, each with one line why.
- [ ] **Problem:** `go test ./...` once you have Chapter 11 tests — come back and check this box then.
