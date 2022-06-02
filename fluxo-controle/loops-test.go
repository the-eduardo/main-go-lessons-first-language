package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 20; i++ {
		if i == 3 {
			// faz isso quando o número não é par
			continue
		}
		fmt.Printf("%v ", i)
	}

  fmt.Println("\n----------------------------------------------------")
  defer fmt.Printf("|")
  secnd()
  //fmt.Printf("|")

}



func secnd() {

	x := 0

	for x <= 21 {

		x++

		if (x % 2 != 0) {
			continue
		}

		fmt.Printf("| %v ", x)


	}

}
