package main

import ("fmt"
)

func multiplier(x, y int32) int32 {
  return x * y
}
// --------------------------- MARS YEARS ---------------------------
func computeMarsYears(earthYears int) int {
  //Computing the value from earth "myAge", thats aewsome!

  earthDays := earthYears * 365
  marsYears := earthDays / 687
  return marsYears
}
//--------------------------- MAIN ---------------------------
func main() {
  var product int32
  product = multiplier(25, 4)
  fmt.Println(product) // prints 100
fmt.Println("****************")

  var main1, main2, newProduct int32
  main1 = 6
  main2 = 7
  newProduct = multiplier(main1, main2)
  fmt.Println(newProduct) // prints 42
  //--------------------------- Main Mars ---------------------------
fmt.Println("****************")
  myAge := 22

// Call marsYear with myAge
myMartianAge := computeMarsYears(myAge)
fmt.Println("Your martian age is ", myMartianAge)
}
