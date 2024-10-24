package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", customer.Name, "my name is", name)
}

func main() {
	rani := Customer{
		Name:    "Rani",
		Address: "Binjai",
		Age:     20,
	}
	rani.sayHello("Fadjrir")
}
