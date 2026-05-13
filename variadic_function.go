package main

import "fmt"

func sumAll(numbers ...int)int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10,10,20,50))
	fmt.Println(sumAll(10,10,20,50,20,10,10))
	fmt.Println(sumAll(10,10,20,50,50,50,20,10,10))

	numbers := []int{10,10,10,10, 10, 10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}