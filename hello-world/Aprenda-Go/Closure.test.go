package main

import "fmt"

func main() {
	a := i()
	fmt.Println("\t", a())
	fmt.Println("\t", a())
	fmt.Println("\t", a())

	v := i()
	fmt.Println("v ", v())
	fmt.Println("v ", v())
	fmt.Println("v ", v())

	fmt.Println("\t", a())
	fmt.Println("\t", a())
	fmt.Println("\t", a())
}
func i() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
