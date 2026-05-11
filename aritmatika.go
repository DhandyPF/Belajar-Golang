package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 30

	// Penjumlahan
	var penjumlahan = a + b + c
	fmt.Println(penjumlahan)

	// Pengurangan
	var pengurangan = c - b - a
	fmt.Println(pengurangan)

	// Perkalian
	var perkalian = a * b * c
	fmt.Println(perkalian)

	// Pembagian
	var pembagian = c / b / a
	fmt.Println(pembagian)

	// Modulus
	var modulus = c % b % a
	fmt.Println(modulus)
}