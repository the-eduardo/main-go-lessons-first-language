package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i + 1)
	var u uint = uint(f + 1)
	fmt.Printf("%v | 	%v  | %v\n", i, f, u)
	fmt.Printf("%T | %T | %T\n", i, f, u)
}
