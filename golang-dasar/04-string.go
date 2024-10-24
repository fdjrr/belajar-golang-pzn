package main

import "fmt"

func main() {
	firstName := "Fadjrir"
	lastName := "Herlambang"
	fullName := firstName + " " + lastName

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)

	fmt.Println("Panjang Nama:", len(fullName))
}
