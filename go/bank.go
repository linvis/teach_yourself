package main

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amout int) {
	mu.Lock()
	balance += amout
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
}
