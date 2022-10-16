package main

import "fmt"

func main() {

	var angka int
	var isOdd bool

	//fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	isOdd = angka%2 != 0

	fmt.Println(isOdd)

}
