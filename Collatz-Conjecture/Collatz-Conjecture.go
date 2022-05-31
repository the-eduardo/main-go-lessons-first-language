// https://www.youtube.com/watch?v=094y1Z2wpJg&t=1098s

//a number is odd *3+1
// if even /2

package main

import "fmt"

func main() {
	var sum uint64
	for sum = 27; sum > 1; {
		if sum&1 != 0 {
			sum = sum*3 + 1
			fmt.Printf("%v = ODD\n", sum)
		} else {
			sum = sum / 2
			fmt.Printf("%v = EVEN\n", sum)
		}
	}
}
