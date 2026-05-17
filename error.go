package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi sama dengan Nol")
	} else	{
		return nilai / pembagi, nil
	}
}

func main() {
	Hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil :", Hasil)
	} else {
		fmt.Println("Error :", err.Error())
	}
}