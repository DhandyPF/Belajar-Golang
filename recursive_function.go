package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	} 
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {

	fmt.Println(factorialRecursive(10))

	// fmt.Println(factorialLoop(1))
	// fmt.Println(factorialLoop(2))
	// fmt.Println(factorialLoop(3))
	// fmt.Println(factorialLoop(4))
	// fmt.Println(factorialLoop(5))
	// fmt.Println(factorialLoop(6))
	// fmt.Println(factorialLoop(7))
	// fmt.Println(factorialLoop(8))
	// fmt.Println(factorialLoop(9))

	fmt.Println(factorialLoop(10))
	
}