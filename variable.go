package main

import "fmt"

func main() {
	var nama string

	nama = "Dhandy Putra Fahruddin"
	fmt.Println(nama)

	name := "Dhandy Putra"
	fmt.Println(name)

	name = "Dhandy Fahruddin"
	fmt.Println(name)

	var (
		firstName = "Dhandy"
		middleName = "Putra"
		lastName = "Fahruddin"
	)

	fmt.Println(firstName + " " + middleName + " " + lastName)

	var umur = 19
	fmt.Println(umur)
}