package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	mychannel := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		mychannel <- true
	}()

	for {
		select {
		case <-mychannel:
			fmt.Println("completed")
			return
		case tm := <-ticker.C:
			fmt.Println(tm)
		}
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time) // the loop doesn't stop until you terminated
	}

}
