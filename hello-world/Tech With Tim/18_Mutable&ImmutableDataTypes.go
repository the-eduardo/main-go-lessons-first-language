package main

import "fmt"

func main() {
	func() {
		var x = 5
		y := &x
		*y = 10
		fmt.Println("With Pointers: x =", x, "| y =", *y)
	}()
	func() {
		var x = 5
		y := x
		y = 10
		fmt.Println("W.No Pointers: x =", x, "| y =", y)
	}()
	fmt.Println("------------------------------------------------------")
	func() {
		var x = []int{1, 2, 3, 4, 5}
		y := x
		fmt.Println("x == y", x, y)
		y[1] = 50
		fmt.Println("y is x", x, y)
	}()
}
