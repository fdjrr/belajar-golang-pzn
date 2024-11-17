package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterWithTimeout(ctx context.Context) chan int {
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

				time.Sleep(1 * time.Second)
			}
		}
	}()

	return dest
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	dest := CreateCounterWithTimeout(ctx)

	for n := range dest {
		fmt.Println("Counter:", n)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}
