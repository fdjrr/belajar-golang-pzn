package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Println(field.Name, "with type", field.Type)
		fmt.Println(field.Tag.Get("required"))
		fmt.Println(field.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()

			result = data != ""

			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{"Fadjrir"})
	readField(Person{"Fadjrir", "Jakarta", "fadjrir@go.dev"})

	person := Person{"Fadjrir", "Jakarta", "fadjrir@go.dev"}
	fmt.Println(isValid(person))
}
