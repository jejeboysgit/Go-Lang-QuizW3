package main

import (
	"fmt"
)

func main() {
	var bil int
	var isZero bool

	// input
	//fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bil)

	// process
	isZero = bil == 0

	// output
	fmt.Println(isZero)
}
