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

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type BankAccount struct {
	mu      sync.Mutex
	balance uint64
}

func (b *BankAccount) Deposite(amount uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *BankAccount) Balance() uint64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.balance
}

func (b *BankAccount) Withdraw(amount uint64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.balance < amount {
		return false
	}

	b.balance -= amount
	return true
}

func main() {
	account := BankAccount{balance: 1000}

	var wg sync.WaitGroup

	n := 20
	for range n {
		wg.Add(1)

		go func() {
			time.Sleep(time.Microsecond * time.Duration(rand.IntN(100)))

			defer wg.Done()

			account.Deposite(100)
		}()
	}
	for range n {
		wg.Add(1)
		go func() {
			time.Sleep(time.Microsecond * time.Duration(rand.IntN(100)))

			defer wg.Done()

			fmt.Printf(
				"Withdraw(150): %v Balance(): %v\n",
				account.Withdraw(150),
				account.Balance())
		}()
	}

	wg.Wait()

	fmt.Printf("account.balance: %v\n", account.balance)

}
