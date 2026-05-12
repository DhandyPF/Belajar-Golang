package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Dimas"
	// person["address"] = "Bukateja"

	person := map[string]string{
		"name" : "Dimas",
		"address" : "Bukateja",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Dimas"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}