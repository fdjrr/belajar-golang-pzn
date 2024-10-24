package main

import "fmt"

/*
Kunci Var

var name string = "Fadjrir Herlambang"
var name = "Fadjrir Herlambang"
name := "Fadjrir Herlambang"

var (
    firstName = "Fadjrir"
    lastName  = "Herlambang"
)

*/

func main() {
	var name string

	name = "Fadjrir Herlambang"
	fmt.Println(name)

	name = "Fadjrir"
	fmt.Println(name)

	var (
		firstName = "Fadjrir"
		lastName  = "Herlambang"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
