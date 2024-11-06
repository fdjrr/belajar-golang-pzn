package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (b *BankAccount) AddBalance(amount int) {
	b.RWMutex.Lock()
	defer b.RWMutex.Unlock()

	b.Balance = b.Balance + amount
}

func (b *BankAccount) GetBalance() int {
	b.RWMutex.RLock()
	defer b.RWMutex.RUnlock()

	balance := b.Balance

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i <= 100; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final balance: ", account.GetBalance())
}
