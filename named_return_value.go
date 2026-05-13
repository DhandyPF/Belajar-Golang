package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Dhandy"
	middleName = "Putra"
	lastName = "Fahruddin"
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}