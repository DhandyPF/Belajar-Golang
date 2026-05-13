package main

import "fmt"

func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Println("Hello", firstName, lastName, age, "Years old")
}

func main()  {
	sayHelloTo("Dhandy", "PF", 19)
}