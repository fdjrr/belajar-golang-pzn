package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Fadjrir"
	lastName = "Herlambang"

	return firstName, lastName
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName, _ = getFullName()
	fmt.Println(firstName)
}
