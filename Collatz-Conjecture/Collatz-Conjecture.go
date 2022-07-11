// https://www.youtube.com/watch?v=094y1Z2wpJg&t=1098s

//a number is odd *3+1
// if even /2

package main

import "fmt"

func main() {
	var sum, max uint64
	for sum = 27; sum > 1; {
		if sum&1 != 0 {
			sum = sum*3 + 1
				if sum > max { // New Feature: Prints the higher value
					max = sum
				}
			fmt.Printf("%v = ODD\n", sum)
		} else {
			sum = sum / 2
			if sum > max {
				max = sum
			}
			fmt.Printf("%v = EVEN\n", sum)
		}
	}
	fmt.Println("\nmax value:",max)
}
