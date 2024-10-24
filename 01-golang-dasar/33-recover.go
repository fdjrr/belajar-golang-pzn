package main

import "fmt"

func endApp() {
	fmt.Println("End Application")

	message := recover()
	fmt.Println("Terjadi Panic:", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Application Error")
	}
}

func main() {
	runApp(true)
}
