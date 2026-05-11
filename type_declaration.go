package main

import "fmt"

func main() {

	type NoHP string

	var NoDimas NoHP = "081234567890"

	var contoh string = "080987654321"
	var contohNoHP NoHP = NoHP(contoh)

	fmt.Println(NoDimas)
	fmt.Println(contohNoHP)
}