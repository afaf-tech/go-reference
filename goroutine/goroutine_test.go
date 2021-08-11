package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // will be run asyncronously with goroutine
	fmt.Println("UPs")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i) // will not be run sequentially
	}

	time.Sleep(5 * time.Second)
}
