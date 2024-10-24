package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang Dasar"

func SayHello(name string) string {
	return "Hello " + name
}

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func Contoh() {
	fmt.Println(sayGoodBye("Herlambang"))
	fmt.Println(version)
}
