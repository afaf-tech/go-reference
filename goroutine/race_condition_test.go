package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceConditions(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		// possible to meet race condition
		/**testing.
		* when (x+1) is assigned to x outside. there is possibility when x meets the same value.
		* this causes some of value assigned to x will be overlapped. and will not generate 100000.
		 */
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Counter  = ", x)
}
