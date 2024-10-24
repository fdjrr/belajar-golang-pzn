package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"fadjrir", "herlambang"})
	_ = writer.Write([]string{"ananda", "maharani"})
	_ = writer.Write([]string{"rani", "herlambang"})

	writer.Flush()
}
