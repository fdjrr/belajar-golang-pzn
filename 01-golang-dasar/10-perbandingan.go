package main

import "fmt"

func main() {
	// Ternary Operator
	fmt.Println("5 > 5 =", 5 > 5)

	// Conditional Operator
	fmt.Println("5 == 5 =", 5 == 5)
	fmt.Println("5 != 5 =", 5 != 5)

	// Logical Operator
	fmt.Println("5 > 5 and 5 == 5 =", 5 > 5 && 5 == 5)
	fmt.Println("5 > 5 or 5 == 5 =", 5 > 5 || 5 == 5)

	// Bitwise Operator
	fmt.Println("5 & 5 =", 5&5)
	fmt.Println("5 | 5 =", 5|5)
	fmt.Println("5 ^ 5 =", 5^5)
}
