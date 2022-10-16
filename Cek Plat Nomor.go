package main

import (
	"fmt"
)

func main() {
	var platNomor string
	var isBandung bool

	// input
	fmt.Print("Masukkan plat nomor: ")
	fmt.Scanln(&platNomor)

	// process
	isBandung = platNomor == "D" || platNomor == "d"

	// output
	fmt.Println(isBandung)
}
