# Learn Go by Coding

I am learning Go by writing programs, not by reading a chapter end to end. This repo follows *The Go Programming Language* (Donovan & Kernighan). The book is a reference I open when a task is stuck.

## How this works with the agent

The agent **generates questions**. I **answer them myself in code**.

1. The agent adds a `.go` file under `NN-chapter/questions/`. The file is a question: a comment block describes what to build, which syntax and ideas to use, and how to run it.
2. I run `make questions`, then fill in the copy in `NN-chapter/myAnswer/`. The agent must not write the solution, complete `TODO`s, or paste the finished implementation.
3. I run the file, fix compiler errors, and check the chapter README box when I am done.
4. If I am stuck, I read **only that section** of the book, then come back and write the code from memory.

Do not copy-paste from the book or the internet. Struggle first.

## Files

The agent writes question templates under each chapter’s `questions/` directory, for example `01-tutorial/questions/1.1-hello-world.go`. Those files are the generated samples. I copy them into `myAnswer/` and edit the copy.

```bash
make questions                 # copy a template into myAnswer/ only if that file is not already there
make clean-unchanged-questions # delete myAnswer files that still match the template; keep edits
make clean-all-questions       # delete every copy in myAnswer/ (templates in questions/ stay)
make watch                     # go run the file whenever I save it under myAnswer/
```

`make questions` never overwrites a file I already started. `make clean-unchanged-questions` only removes copies I have not changed. `make clean-all-questions` removes all working copies, including edits.

For a fast edit loop, copy the questions, then leave watch running in a terminal. Saving a `.go` file in any chapter `myAnswer/` runs that file (`go run`, or `go test` for `*_test.go`). Saving again stops the previous run first, so a web server can be restarted by saving.

```bash
make questions
make watch
```

Or run one file by hand:

```bash
cd 01-tutorial/myAnswer
go run 1.1-hello-world.go
```

When I want a module at the repo root:

```bash
go mod init github.com/abolfazl/learn-go
```



## Chapters


| Chapter                              | Folder                                                                                    |
| ------------------------------------ | ----------------------------------------------------------------------------------------- |
| 1. Tutorial                          | `[01-tutorial/](01-tutorial/README.md)`                                                   |
| 2. Program Structure                 | `[02-program-structure/](02-program-structure/README.md)`                                 |
| 3. Basic Data Types                  | `[03-basic-data-types/](03-basic-data-types/README.md)`                                   |
| 4. Composite Types                   | `[04-composite-types/](04-composite-types/README.md)`                                     |
| 5. Functions                         | `[05-functions/](05-functions/README.md)`                                                 |
| 6. Methods                           | `[06-methods/](06-methods/README.md)`                                                     |
| 7. Interfaces                        | `[07-interfaces/](07-interfaces/README.md)`                                               |
| 8. Goroutines and Channels           | `[08-goroutines-and-channels/](08-goroutines-and-channels/README.md)`                     |
| 9. Concurrency with Shared Variables | `[09-concurrency-with-shared-variables/](09-concurrency-with-shared-variables/README.md)` |
| 10. Packages and the Go Tool         | `[10-packages-and-the-go-tool/](10-packages-and-the-go-tool/README.md)`                   |
| 11. Testing                          | `[11-testing/](11-testing/README.md)`                                                     |
| 12. Reflection                       | `[12-reflection/](12-reflection/README.md)`                                               |
| 13. Low-Level Programming            | `[13-low-level-programming/](13-low-level-programming/README.md)`                         |


Preface pages (origins, the Go project, how the book is organized): skim once, then start Chapter 1.

## Progress

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

- Compile often. Read the error. Fix that error.
- Break programs on purpose. The compiler is part of the lesson.
- Run `gofmt` / `go fmt` on a file before leaving it.
- Chapter README checklists are the map. `.go` question files are the work.
- The agent generates the next question when I ask. I write every solution.

