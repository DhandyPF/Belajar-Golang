package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error{
		panic("Ups error")
	}
}

func main() {
	runApp(false)
}