package main

import "fmt"

func getFullName() (string, string) {
	return "Fadjrir", "Herlambang"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName, _ = getFullName()
	fmt.Println(firstName)
}
