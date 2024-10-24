package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var rani Customer
	rani.Name = "Rani"
	rani.Address = "Binjai"
	rani.Age = 20
	fmt.Println(rani)

	fadjrir := Customer{
		Name:    "Fadjrir",
		Address: "Medan",
		Age:     19,
	}
	fmt.Println(fadjrir)

	herlambang := Customer{"Herlambang", "Jakarta", 21}
	fmt.Println(herlambang)
}
