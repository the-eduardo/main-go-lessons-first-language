package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x), len(y))
  fmt.Printf("\n-----------------------------------------------\n")
  arr()
}

//        https://gobyexample.com/arrays
func arr() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("b value:", b)
    fmt.Printf("-----------------------------------------------")
    var twoD [2][3]int
    fmt.Println("\nigual:", twoD)
    for i := 0; i < 2; i++ {
      fmt.Printf("\n----------")
        for j := 0; j < 3; j++ {
                    twoD[i][j] = i + j
                    fmt.Println("2d mid: ", twoD, i + j)
        }
    }
    fmt.Println("\t2d: ", twoD)
}
