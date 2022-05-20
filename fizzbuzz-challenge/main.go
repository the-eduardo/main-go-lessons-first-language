package main

import (
	"fmt"
)

func fizz(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}

func main() {

	n := 1
	for n < 101 {
		fizz(n)
		n += 1
	}

}
