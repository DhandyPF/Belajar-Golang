package main

import "fmt"

type customer struct {
	Name, Address string
	Age int
}

func main() {
	var a customer
	a.Name = "Dimas,"
	a.Address = "Depok,"
	a.Age = 100
	fmt.Println(a)

	b := customer {
		Name: "Willy,",
		Address: "Sokaraja,",
		Age: 11,
	}
	fmt.Println(b)

	c := customer{"Kautsar,", "PWT,", 10}
	fmt.Println(c)
}