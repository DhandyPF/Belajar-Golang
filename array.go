package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Dhandy"
	name[1] = "Putra"
	name[2] = "Fahruddin"
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{10, 3, 2007}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var values2 = [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
	}

	fmt.Println(values2)

	fmt.Println(len(values))
	fmt.Println(len(values2))
}