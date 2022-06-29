package main

import "fmt"

func main() {
	var x float64 = 4
	fmt.Println(fatorial(x))
	fmt.Println(loops(x))
}
func fatorial(x float64) float64 {
	if x <= 1 {
		return x
	}
	return x * fatorial(x-1)
}
func loops(x float64) float64 {
	total := 1.0
	for x > 1 {
		total *= x
		x--
	}
	return total
}
