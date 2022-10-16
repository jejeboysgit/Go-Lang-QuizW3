package main

import (
	"fmt"
)

func main() {
	var bilAdik, bilKakak int
	var isAdikWin bool

	// input
	//fmt.Print("Adik: ")
	//fmt.Scanln(&bilAdik)
	//fmt.Print("Kakak: ")
	//fmt.Scanln(&bilKakak)
  
  fmt.Scanln(&bilAdik, &bilKakak)

	// process
	isAdikWin = bilAdik == bilKakak || bilAdik == bilKakak*3

	// output
	fmt.Println(isAdikWin)
}
