package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch result.(type) {
	case string:
		fmt.Println("String", result.(string))
	case int:
		fmt.Println("Integer", result.(int))
	default:
		fmt.Println("Unknown Type")
	}
}
