package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.4))  // bulat ke atas
	fmt.Println(math.Round(1.7)) // bulat yg terdekat
	fmt.Println(math.Floor(1.4)) // bulat ke bawah
	fmt.Println(math.Max(1.4, 1.5))
	fmt.Println(math.Min(1.4, 1.5))
}
