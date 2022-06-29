package main

import "fmt"

func main() {
	t := somentePares(soma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 50, 100, 101, 103}...)
	fmt.Println(t)
}
func soma(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	fmt.Println(slice)
	total := f(slice...)
	return total
}
