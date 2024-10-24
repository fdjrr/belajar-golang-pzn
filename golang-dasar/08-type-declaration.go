package main

import "fmt"

func main() {
	type NoKTP string

	var noKTP1 NoKTP = "1234567890"

	var contoh string = "1234567890"
	var noKTP2 NoKTP = NoKTP(contoh)

	fmt.Println(noKTP1)
	fmt.Println(noKTP2)
}
