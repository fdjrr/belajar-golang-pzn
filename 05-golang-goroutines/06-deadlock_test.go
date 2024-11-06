package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Name    string
	Balance int
	sync.Mutex
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance = u.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	defer user1.Unlock()
	user1.Change(-amount)
	fmt.Println("Lock User1 :", user1.Name)

	time.Sleep(1 * time.Second)

	user2.Lock()
	defer user2.Unlock()
	user2.Change(amount)
	fmt.Println("Lock User2 :", user2.Name)

	time.Sleep(1 * time.Second)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Herlambang",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Fadjrir",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User1 :", user1.Name, "Balance :", user1.Balance)
	fmt.Println("User2 :", user2.Name, "Balance :", user2.Balance)
}
