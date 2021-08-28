package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	if value%2 == 0 {
		cond.Wait() // create some condition and wait
	}
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // signal is used for stopping cond.Wait no need to wait anymore.
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() // ingnore all cond.Wait
	// }()

	group.Wait()
}
