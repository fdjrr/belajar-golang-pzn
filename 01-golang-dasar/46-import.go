package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	result := helper.SayHello("Golang")

	fmt.Println(result)

	fmt.Println(helper.Application)
}
