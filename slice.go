package main

import "fmt"

func main() {
	moons := [...]string{
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

	// fmt.Println(moons)

	slice1 := moons[5:10]
	fmt.Println(slice1)

	slice2 := moons[:3]
	fmt.Println(slice2)

	slice3 := moons[9:]
	fmt.Println(slice3)

	slice4 := moons[:]
	fmt.Println(slice4)

	panjangSlice1 := len(slice1)
	fmt.Println(panjangSlice1)

	kapasitasSlice1 := cap(slice1)
	fmt.Println(kapasitasSlice1)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
}