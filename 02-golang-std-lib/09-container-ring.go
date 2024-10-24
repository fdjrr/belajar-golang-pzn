package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = i
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
