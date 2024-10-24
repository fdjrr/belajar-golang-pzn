package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err)
	}

	resultInt, err := strconv.Atoi("10000")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println(err)
	}

	binary := strconv.FormatInt(1000, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(1000)
	fmt.Println(stringInt)
}
