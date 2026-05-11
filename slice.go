package main

import "fmt"

func main() {
	bulan := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// fmt.Println(bulan)

	slice1 := bulan[5:10]
	fmt.Println(slice1)

	slice2 := bulan[:3]
	fmt.Println(slice2)

	slice3 := bulan[9:]
	fmt.Println(slice3)

	slice4 := bulan[:]
	fmt.Println(slice4)
}