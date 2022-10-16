package main

import "fmt"

func main() {

  var isMember, hasCashback bool
  var totalBelanja int64

  //fmt.Print("Apakah member? ")
  //fmt.Scanln(&isMember);

  //fmt.Print("Total Belanja: ")
  //fmt.Scanln(&totalBelanja);
  
  fmt.Scanln(&isMember, &totalBelanja);

  hasCashback = isMember && totalBelanja >= 1000000

  fmt.Println("\n Apakah mendapatkan cashback ? ", hasCashback)
}
