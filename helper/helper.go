package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

func SayHello(name string) string {
	return "Hello " + name
}

func sayGoodbye(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodbye("Dimas")
	fmt.Println(version)
}