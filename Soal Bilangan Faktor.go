package main

import "fmt"

/**
* faktor bilangan adalah sebuah bilangan yang dapat 
* membagi habis bilangan tersebut
*/
func main() {
  var bil0, bil1 int
	var isFactor bool

	//fmt.Print("Masukkan bilangan ke-1: ")
	//fmt.Scanln(&bil0)
  //fmt.Print("Masukkan bilangan ke-2: ")
	//fmt.Scanln(&bil1)
  
  fmt.Scanln(&bil0, &bil1)

	isFactor = bil1 % bil0 == 0

	fmt.Println(isFactor)
}
