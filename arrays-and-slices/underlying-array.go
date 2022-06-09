package main

import "fmt"

func main() {
	fslice := []int{1, 2, 3, 4, 5}
	fmt.Println(fslice)
	sslice := append(fslice[:2], fslice[4:]...)
	fmt.Println(sslice)
	fmt.Println(fslice)
}

/*  Output:
[1 2 3 4 5]
[1 2 5]
[1 2 5 4 5]
*/
