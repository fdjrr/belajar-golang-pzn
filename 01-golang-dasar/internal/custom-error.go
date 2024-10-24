package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{
			Message: "ID cannot be empty",
		}
	}

	if id != "fadjrir" {
		return &notFoundError{
			Message: "Data not found",
		}
	}

	return nil
}

func main() {
	// err := SaveData("", nil)
	err := SaveData("fadjrir", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println(validationErr.Message)
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println(notFoundErr.Message)
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println(finalError.Message)
		case *notFoundError:
			fmt.Println(finalError.Message)
		default:
			fmt.Println("Data Not Saved")
		}
	} else {
		fmt.Println("Data Saved")
	}
}
