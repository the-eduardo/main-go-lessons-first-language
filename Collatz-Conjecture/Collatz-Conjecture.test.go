// https://www.youtube.com/watch?v=094y1Z2wpJg&t=1098s

//a number is odd *3+1
// if even /2

package main

import "fmt"

func main() {
	o, e := 0, 0
	var sum uint64
	_, err := fmt.Scan(&sum)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for sum > 1 {
		if sum&1 != 0 {
			o++
			sum = sum*3 + 1
			fmt.Printf("%v = ODD\n", sum)
		} else {
			e++
			sum = sum / 2
			fmt.Printf("%v = EVEN\n", sum)
		}
	}
	fmt.Printf("\n Evens: %v | Odds: %v | Total: (%v)", e, o, e+o)

}
