package main

import "fmt"

type customer struct {
	Name, Address string
	Age int
}

func main() {
	var dimas customer
	dimas.Name = "Dimas,"
	dimas.Address = "Depok,"
	dimas.Age = 100
	fmt.Println(dimas)

	willy := customer {
		Name: "Willy,",
		Address: "Sokaraja,",
		Age: 11,
	}
	fmt.Println(willy)

	kautsar := customer{"Kautsar,", "PWT,", 10}
	fmt.Println(kautsar)
}