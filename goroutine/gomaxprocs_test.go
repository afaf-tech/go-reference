package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(time.Second * 3)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Printf("Total CPU: %d\n", totalCpu)

	runtime.GOMAXPROCS(20)                 // change cpu thread
	totalThreads := runtime.GOMAXPROCS(-1) // if just view then set value below 0.
	fmt.Printf("Total Threads: %d\n", totalThreads)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Printf("Totalgoroutine: %d\n", totalGoroutine)

	group.Wait()
}
