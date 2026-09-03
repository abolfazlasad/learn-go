// Question 9.2 — sync.Mutex
//
// type Bank struct { mu sync.Mutex; balance int }
// Deposit(amount int), Balance() int, Withdraw(amount int) bool
// Withdraw must check and subtract in one critical section.
//
// Hammer Deposit/Withdraw from several goroutines. Final balance must be consistent.
//
// You will need: sync.Mutex, Lock, defer Unlock.
//
// Run:
//   go run -race 9.2-bank.go

package main

func main() {
	// TODO: write your solution
}
