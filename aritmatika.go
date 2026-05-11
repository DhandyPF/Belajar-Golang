package main

import "fmt"

func main() {
	var a = 20
	var b = 30
	var c = 40

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