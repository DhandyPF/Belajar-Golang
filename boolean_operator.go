package main

import "fmt"
func main() {
	// Boolean Operator
	var nilaiAkhir = 90
	var absensi = 70

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	var lulus2 bool = lulusNilaiAkhir || lulusAbsensi
	fmt.Println(lulus)
	fmt.Println(lulus2)
}