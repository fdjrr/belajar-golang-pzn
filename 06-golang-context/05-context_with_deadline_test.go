package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterWithDeadline(ctx context.Context) chan int {
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

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	dest := CreateCounterWithDeadline(ctx)

	for n := range dest {
		fmt.Println("Counter:", n)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}
