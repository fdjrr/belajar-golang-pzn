package main

import "fmt"

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	result := 5 * 4 * 3 * 2 * 1
	fmt.Println(result)

	fmt.Println(factorialLoop(5))

	fmt.Println(factorialRecursive(5))
}
