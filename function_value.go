package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye " + name 
}

func main() {
	contoh := getGoodbye
	misal := getGoodbye

	fmt.Println(contoh("Dimas"))
	fmt.Println(misal("Dimas"))
}