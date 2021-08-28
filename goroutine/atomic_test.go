package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0

	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		// possible to meet race condition
		/**testing.
		* when (x+1) is assigned to x outside. there is possibility when x meets the same value.
		* this causes some of value assigned to x will be overlapped. and will not generate 100000.
		 */
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				// x = x + 1 changed to the next line
				atomic.AddInt64(&x, 1) // prevents race condition
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter  = ", x)
}
