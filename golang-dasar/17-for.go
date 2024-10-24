package main

import "fmt"

func main() {
	count := 1

	for count <= 10 {
		fmt.Println("Perulangan Ke", count)
		count++
	}

	count = 10

	for i := 0; i < count; i++ {
		fmt.Println("Hello World")
	}

	names := []string{"Fadjrir", "Herlambang"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
