package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fatih Al Fikri" // send data to channel
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "Ahmad Akmal Fikri"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	// data := <-channel // invalid operation
	channel <- "Ahmad Akmal Fikri"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)

	go OnlyOut(channel)

}

/**
* buffer channel is used for delegate more than one data in a channel
* defalt buffer when create channel is 1.
 */
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // 3= length of buffer
	defer close(channel)

	go func() {
		channel <- "Fikri"
		channel <- "Al"
		channel <- "Fatih"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {

	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// used to get unkown length of data from channel
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}
