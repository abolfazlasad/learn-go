# Learn Go by Coding

This repo follows *The Go Programming Language* (Donovan & Kernighan). You learn each topic by writing programs and solving problems, not by reading a chapter end to end.

The book is a reference you open when a task is stuck. Close it again before you write the next program.

## How to work a section

1. Open the chapter folder. Put that section's program there, named like `1.1-hello-world.go`.
2. Try the first coding task **without** reading the section.
3. When you are stuck, read **only that section** in the book. Type examples yourself. Do not copy-paste.
4. Close the book. Finish the drills from memory, then the problems.
5. Check the boxes in that chapter's `README.md`.
6. Do not start the next section until you can explain the concept out loud and the problems pass.

One file per idea. Run a single file so programs do not collide:

```bash
cd 01-tutorial
go run 1.1-hello-world.go
```

When you want a module at the repo root:

```bash
go mod init github.com/abolfazl/learn-go
```



## Chapters


| Chapter | Folder |
| --- | --- |
| 1. Tutorial | [`01-tutorial/`](01-tutorial/README.md) |
| 2. Program Structure | [`02-program-structure/`](02-program-structure/README.md) |
| 3. Basic Data Types | [`03-basic-data-types/`](03-basic-data-types/README.md) |
| 4. Composite Types | [`04-composite-types/`](04-composite-types/README.md) |
| 5. Functions | [`05-functions/`](05-functions/README.md) |
| 6. Methods | [`06-methods/`](06-methods/README.md) |
| 7. Interfaces | [`07-interfaces/`](07-interfaces/README.md) |
| 8. Goroutines and Channels | [`08-goroutines-and-channels/`](08-goroutines-and-channels/README.md) |
| 9. Concurrency with Shared Variables | [`09-concurrency-with-shared-variables/`](09-concurrency-with-shared-variables/README.md) |
| 10. Packages and the Go Tool | [`10-packages-and-the-go-tool/`](10-packages-and-the-go-tool/README.md) |
| 11. Testing | [`11-testing/`](11-testing/README.md) |
| 12. Reflection | [`12-reflection/`](12-reflection/README.md) |
| 13. Low-Level Programming | [`13-low-level-programming/`](13-low-level-programming/README.md) |


Preface pages (origins, the Go project, how the book is organized): skim once, then start Chapter 1. No folder.

## Progress

Mark these when a whole chapter is done.

- [ ] 1. Tutorial
- [ ] 2. Program Structure
- [ ] 3. Basic Data Types
- [ ] 4. Composite Types
- [ ] 5. Functions
- [ ] 6. Methods
- [ ] 7. Interfaces
- [ ] 8. Goroutines and Channels
- [ ] 9. Concurrency with Shared Variables
- [ ] 10. Packages and the Go Tool
- [ ] 11. Testing
- [ ] 12. Reflection
- [ ] 13. Low-Level Programming



## Rules

- Compile often. Read the error. Fix that error. Do not guess in the dark.
- Break programs on purpose. The compiler is part of the lesson.
- Prefer `gofmt` / `go fmt .` before you leave a folder.
- Book exercises at the end of a section are extra, not a replacement for the problems here.
- Do not copy solutions from the internet. Struggle, then look only at the book section.

