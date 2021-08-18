package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex        // Protect
	Name       string //
	balance    int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.balance = user.balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
}

func TransferWithWaitGroup(group *sync.WaitGroup, user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)
	user1.Unlock()

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	user2.Unlock()
	defer group.Done()
}

// deadlock = kejadian dua goroutine sama-sama menunggu routine to unlock
func TestTransferDeadLock(t *testing.T) {
	user1 := UserBalance{Name: "Fikri", balance: 1000000}
	user2 := UserBalance{Name: "Budi", balance: 1000000}

	// deadlock happens
	/**
	* before the user2 executed, it will be locked by the second go routine
	 */
	go Transfer(&user1, &user2, 1000) // transfer from user 1 to user 2
	go Transfer(&user2, &user1, 1000) // transfer from user 2 to user 1

	time.Sleep(3 * time.Second) // if not exist then the next code will just print the current value
	fmt.Println("User ", user1.Name, ", Balance ", user1.balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.balance)
}

func TestTransferWithWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	user1 := UserBalance{Name: "Fikri", balance: 1000000}
	user2 := UserBalance{Name: "Budi", balance: 1000000}

	// deadlock happens
	/**
	* before the user2 executed, it will be locked by the second go routine
	 */
	group.Add(1)
	go TransferWithWaitGroup(group, &user1, &user2, 60000) // transfer from user 1 to user 2
	group.Wait()
	group.Add(1)
	go TransferWithWaitGroup(group, &user2, &user1, 1000) // transfer from user 2 to user 1
	group.Wait()

	fmt.Println("User ", user1.Name, ", Balance ", user1.balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.balance)
}
