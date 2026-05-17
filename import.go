package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Dimas")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodbye("Dimas"))
}