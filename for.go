package main

import "fmt"

func main() {
	// count := 1

	// for count <= 10 {
	// 	fmt.Println("Perulangan ke : ", count)
	// 	count++		
	// }
	// fmt.Println("Selesai")

	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan ke : ", count)
	}
	fmt.Println("Selesai")

	names := []string {"Dhandy", "Putra", "Fahruddin"}
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	for index, name := range names{
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names{
		fmt.Println(name)
	}
}