package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterWithCancel(ctx context.Context) chan int {
	dest := make(chan int)

	go func() {
		defer close(dest)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- counter
				counter++
			}
		}
	}()

	return dest
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	dest := CreateCounterWithCancel(ctx)

	for n := range dest {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(3 * time.Second)

	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}
