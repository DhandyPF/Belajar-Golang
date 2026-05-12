package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("Perulangan ke : ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke : ", i)
	}
}