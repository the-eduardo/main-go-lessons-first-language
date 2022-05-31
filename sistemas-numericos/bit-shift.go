package main

import "fmt"
const (
  a = iota + 1000
  b
  _
  c
  d
  _
  x
)

func main() {
  fmt.Printf("%v | %v | %v | %v | %v\n", a, b, c, d, x)
  fmt.Println("-----------------------------")
y := 24
z := y
  for z < 2050 {
    z += y >> 1 // didn't worked :V
  fmt.Printf("\n %v | %b | %b | %v", z, z, y, y)
      }
      bitShift()

}

func bitShift(){ // working bitShift
  fmt.Printf("\n\n-----------bitShift--------------")

  aa := 24
  bb := aa << 2
  cc := aa >> 2
  fmt.Printf("\n %v | %b \n %v | %b \n %v  | %b\n", aa, aa, bb, bb, cc, cc)
}
