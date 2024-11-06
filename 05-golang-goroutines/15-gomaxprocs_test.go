package golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGoMaxProcs(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Hello World")
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutines)
}
