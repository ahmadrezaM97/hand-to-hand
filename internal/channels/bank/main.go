package bank

import (
	"fmt"
	"sync"
)

type Account struct {
	Value int
	lock  sync.Mutex
}

func (account *Account) Increment() {
	account.lock.Lock()
	account.Value += 10
	account.lock.Unlock()
}

func (account *Account) Decrement() {
	account.lock.Lock()
	account.Value -= 10
	account.lock.Unlock()
}

func (acc *Account) Print() {
	fmt.Println(acc.Value)
}
