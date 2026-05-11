package main

import "fmt"

func main() {
	var a = 20

	// Penjumlahan
	a += 100
	fmt.Println(a)

	// Penjumlahan dengan dirinya sendiri
	a += a
	fmt.Println(a)

	// Pengurangan
	a -= 5
	fmt.Println(a)

	// Pengurangan dengan dirinya sendiri
	a -= a
	fmt.Println(a)

	// Perkalian
	a *= 2
	fmt.Println(a)

	// Perkalian dengan dirinya sendiri
	a *= a
	fmt.Println(a)

	// Pembagian
	a /= 4
	fmt.Println(a)

	// Pembagian dengan dirinya sendiri
	a /= a
	fmt.Println(a)

	// Modulus
	a %= 3
	fmt.Println(a)

	// Modulus dengan dirinya sendiri
	a %= a
	fmt.Println(a)
}