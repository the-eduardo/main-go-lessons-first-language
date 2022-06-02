package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Printf("%d\n", x)
		for j := 1; j <= 3; j++ {
			fmt.Printf(" %#U\n", x)
		}
	}
}
