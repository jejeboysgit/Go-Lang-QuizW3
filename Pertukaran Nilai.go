package main

import (
	"fmt"
)

func main() {

	var a, b, c, d int

	//fmt.Print("A: ")
	//fmt.Scanln(&a)
	//fmt.Print("B: ")
	//fmt.Scanln(&b)
	//fmt.Print("C: ")
	//fmt.Scanln(&c)
	//fmt.Print("D: ")
	//fmt.Scanln(&d)

  fmt.Scanln(&a, &b, &c, &d)


  d = d - c
  c = c + a
  b = b + a
  a = a + a

	fmt.Println(a, b, c, d)
 
}
