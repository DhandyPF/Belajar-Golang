package main

import "fmt"

func main() {
	name := "Her"

	switch name {
	case "Dhandy":
		fmt.Println("Hello Dhandy")
	case "Dimas":
		fmt.Println("Hello Dimas")
	case "Her":
		fmt.Println("Hello Her")
	default:
		fmt.Println("Hai, siapa kamu?")
	}

	switch length := len(name); length >= 10 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Nama terlalu pendek")
	}
}