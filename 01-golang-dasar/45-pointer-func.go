package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fadjrir := Man{"Fadjrir"}
	fadjrir.Married()

	fmt.Println(fadjrir)
}
