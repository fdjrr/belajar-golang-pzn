package main

import (
	"fmt"
	"strings"
)

func main() {
	fullName := "Fadjrir Herlambang    "

	fmt.Println(fullName)
	fmt.Println(strings.Contains(fullName, "Herlambang"))
	fmt.Println(strings.Split(fullName, " "))
	fmt.Println(strings.ToUpper(fullName))
	fmt.Println(strings.ToLower(fullName))
	fmt.Println(strings.Trim(fullName, " "))
}
