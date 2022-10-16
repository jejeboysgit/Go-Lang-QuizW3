package main

import "fmt"

func main() {

	var bilAdik, bilKakak, selisih int8
	var isPrime bool

	//fmt.Print("Adik: ")
	//fmt.Scanln(&bilAdik)
	//fmt.Print("Kakak: ")
	//fmt.Scanln(&bilKakak)
  
  fmt.Scanln(&bilAdik, &bilKakak)

	selisih = bilAdik - bilKakak
	isPrime = ((selisih/2 == 1 || selisih/2 == -1) && selisih%selisih == 0)

	fmt.Println(isPrime)

}
