package main

import "fmt"

func main() {
	number := 10

	if number > 5 {
		fmt.Println("Lebih besar dari 5")
	} else if number == 5 {
		fmt.Println("Sama dengan 5")
	} else {
		fmt.Println("Kurang dari 5")
	}

	name := "Dimas"

	if length := len(name); length > 10  {
		fmt.Println("Nama terlalu panjang")
	} else if length < 5 {
		fmt.Println("Nama terlalu pendek")
	} else if name == "Dimas" {
		fmt.Println("Dimas ngaku-ngaku Depok")
	} else {
		fmt.Println("Namanya bagus")
	}
}