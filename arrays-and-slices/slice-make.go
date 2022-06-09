package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4
	slice[4] = 10
	fmt.Printf("\n%v | Len: %v | Cap: %v", slice, len(slice), cap(slice))
	slice = append(slice, 11)
	slice = append(slice, 12)
	slice = append(slice, 13)
	slice = append(slice, 14)
	slice = append(slice, 15)
	fmt.Printf("\n%v | Len: %v | Cap: %v", slice, len(slice), cap(slice))
	slice = append(slice, 20)
	fmt.Printf("\n%v | Len: %v | Cap: %v", slice, len(slice), cap(slice))
}
