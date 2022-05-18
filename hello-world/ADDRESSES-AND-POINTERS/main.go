package main

import "fmt"

func addHundred (numPtr *int) {
  *numPtr += 100
}

func main() {
  x := 1
  addHundred(&x)
  fmt.Println(x)
  fmt.Println("****************")

  main2()
}

/* // TODO:
Create a function that changes a boolean value from a different scope.
Create a function that changes a float value from a different scope.
Find the zero value of a pointer that is initialized but hasn’t been assigned a value.  */

func brainwash(saying *string) {
	// Dereference:
	*saying = "Beep Boop."
}

func main2() {
	greeting := "Hello there!"
  fmt.Println(greeting)
  fmt.Println("***brainwash working***")
	// Call:
	brainwash(&greeting)

	fmt.Println("greeting is now:", greeting)
}
