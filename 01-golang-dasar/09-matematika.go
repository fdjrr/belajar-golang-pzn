package main

import "fmt"

func main() {
	fmt.Println("5 + 5 =", 5+5)
	fmt.Println("5 - 5 =", 5-5)
	fmt.Println("5 * 5 =", 5*5)
	fmt.Println("5 / 5 =", 5/5)
	fmt.Println("5 % 5 =", 5%5)

	// Augmented Assignment
	i := 10
	i += 5
	fmt.Println("i =", i)

	// Unary Operator
	i++
	fmt.Println("i =", i)
}
