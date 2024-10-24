package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// hasil, err := Pembagian(10, 0)
	hasil, err := Pembagian(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hasil)
	}
}
