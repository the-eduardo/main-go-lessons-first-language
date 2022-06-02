package main

import "fmt"

func main() {
	x := 0
	switch {
	case x > 10, x == 1:
		fmt.Println("x > 10 or x = 1")
	case x == 10:
		fmt.Println("x é 10")
		fallthrough
	case (x < 10):
		fmt.Println("x < 10 ou é 10")
	default:
		fmt.Println("fon")

	}
}
