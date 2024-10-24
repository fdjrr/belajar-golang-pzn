package main

import "fmt"

func main() {
	name := "Fadjrir"

	if name == "Herlambang" {
		fmt.Println("Hello Herlambang")
	} else if name == "Fadjrir" {
		fmt.Println("Hi, Fadjrir")
	} else {
		fmt.Println("Hai, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
