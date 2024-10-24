package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "fadjrir,herlambang\n" +
		"ananda,maharani\n" +
		"rani,herlambang\n" +
		"nur,herlambang"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
