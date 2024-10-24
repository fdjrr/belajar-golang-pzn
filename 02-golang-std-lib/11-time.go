package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2024-10-24 12:00:00"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)

		fmt.Println(valueTime.Year())
		fmt.Println(valueTime.Month())
		fmt.Println(valueTime.Day())
		fmt.Println(valueTime.Hour())
	}
}
