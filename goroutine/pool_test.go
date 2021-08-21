package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} { // to replace nil
			return "New"
		},
	}

	pool.Put("Fikri")
	pool.Put("Fatih")
	pool.Put("Al")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}
