package main

import "fmt"

func main() {

	// Relational
	var a = 10
	var b = 20

	var result3 bool = a > b
	fmt.Println(result3)

	var result4 bool = a < b
	fmt.Println(result4)

	var result5 bool = a >= b
	fmt.Println(result5)

	var result6 bool = a <= b
	fmt.Println(result6)

	var name1 = "Dhandy"
	var name2 = "Dhandy"

	var result bool = name1 == name2
	fmt.Println(result)

	var result2 bool = name1 != name2
	fmt.Println(result2)
}