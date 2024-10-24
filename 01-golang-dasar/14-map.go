package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fadjrir",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["name"] = "Fadjrir Herlambang"

	fmt.Println(person)

	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Fadjrir Herlambang"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
