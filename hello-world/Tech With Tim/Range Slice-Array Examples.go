package main

import "fmt"

func main() {
	var a []int = []int{11, 1, 2, 40, 41, 578, 1}
	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && i > j {
				fmt.Println(element, i, j)
			}
		}
	}
	fmt.Println("------------------------------------------------------")
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element == element2 {
				fmt.Println(element)
			}
		}
	}
}
