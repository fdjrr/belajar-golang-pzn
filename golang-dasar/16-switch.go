package main

import "fmt"

func main() {
	name := "Rani"

	switch name {
	case "Fadjrir":
		fmt.Println("Hello Fadjrir")
	case "Herlambang":
		fmt.Println("Hello Herlambang")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	default:
		fmt.Println("Sudah Benar")
	}
}
