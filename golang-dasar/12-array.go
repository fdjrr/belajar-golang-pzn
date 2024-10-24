package main

import "fmt"

func main() {
	names := [2]string{}

	names[0] = "Fadjrir"
	names[1] = "Herlambang"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	values := [...]int{1, 2, 3, 4, 5}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[4])

	fmt.Println("Panjang Array:", len(values))
	values[2] = 12

	fmt.Println(values)
}
