package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "fadjrir" {
		return NotFoundError
	}

	return nil
}

func main() {
	// err := GetById("")
	err := GetById("fadjrir")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("ID cannot be empty")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Data not found")
		} else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("Data found")
	}
}
