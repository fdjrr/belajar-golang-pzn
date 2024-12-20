package golang_goroutines

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

		channel <- "Fadjrir Herlambang"

		fmt.Println("Selesai")
	}()

	msg := <-channel

	fmt.Println(msg)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "Fadjrir Herlambang"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go GiveMeResponse(channel)

	msg := <-channel

	fmt.Println(msg)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "Fadjrir Herlambang"
}

func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)

	data := <-channel

	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)

	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	defer close(channel)

	go func() {
		channel <- "Fadjrir"
		channel <- "Herlambang"
		channel <- "Rani"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	counter := 0

	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1:", data)

			counter++

		case data := <-channel2:
			fmt.Println("Channel 2:", data)
			counter++

		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
