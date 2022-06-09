package main

import "fmt"

func main() {
	robson := [5]int{0, 1, 2, 3, 4}
	fmt.Println(robson)
	fmt.Println("\n-----------------------------------------------")
	for i := 0; i < len(robson); i++ {
		fmt.Println(robson[i])
	}
	fmt.Printf("robson é %T", robson)
	for i, v := range robson {
		fmt.Println(i, v)
	}
	fmt.Println("\n-----------------------------------------------")
	slice()
}

func slice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ddd := append(slice, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	fmt.Printf("%v %v", ddd, "\n")
	fmt.Printf("ddd é %T \n", ddd)
	for i, v := range ddd {
		fmt.Printf("%v-%v|", i, v)
	}
	fmt.Println("\n-----------------------------------------------")
	fmt.Println(ddd[:3])
	fmt.Println(ddd[4:])
	fmt.Println(ddd[1:7])
	fmt.Println(ddd[2:19])
	fmt.Println(ddd[2 : len(ddd)-1])
	fmt.Println(ddd[2 : len(ddd)-2])
}
