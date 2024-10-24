package main

import (
	"flag"
	"fmt"
)

/*
go run ./04-flag.go

or

go run ./04-flag.go -username=fadjrir -port=1020
*/

func main() {
	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", 3306, "port")
	username := flag.String("username", "root", "username")
	password := flag.String("password", "root", "password")

	flag.Parse()

	fmt.Println(*host, *port, *username, *password)
}
