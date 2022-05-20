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
	for n <= 100 {
		fizz(n)
		n += 1
	}
	fmt.Println("****************")
second()
}


func second() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println("^ final")
	forContinued()
}


func forContinued() {
	sum := 1
	for ; sum < 1025; { // The init and post statements are optional.
		sum += sum
	}
	fmt.Println(sum)
}



/*-------------- For
	Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

   -- The init statement: executed before the first iteration

   -- The condition expression: evaluated before every iteration

   -- The post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

********************************/
