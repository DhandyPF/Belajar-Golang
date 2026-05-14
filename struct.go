package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func(customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var dimas Customer
	dimas.Name = "Dimas,"
	dimas.Address = "Depok,"
	dimas.Age = 100
	fmt.Println(dimas)

	willy := Customer {
		Name: "Willy,",
		Address: "Sokaraja,",
		Age: 11,
	}
	fmt.Println(willy)

	kautsar := Customer{"Kautsar,", "PWT,", 10}
	fmt.Println(kautsar)

	bintang := Customer{"Bintang,", "Susukan,", 22}
	bintang.sayHello("Bintang")
}