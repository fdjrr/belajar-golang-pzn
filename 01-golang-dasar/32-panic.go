package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("App Error")
	}
}

func main() {
	runApp(true)
}
