package main

import (
	"fmt"
)

func main() {
	var bilangan int
	var hasil bool

	// input
	//fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	// process
	hasil = bilangan == 0

	// output
	fmt.Println(hasil)
}
