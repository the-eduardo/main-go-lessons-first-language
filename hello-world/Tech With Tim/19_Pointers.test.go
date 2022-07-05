package main

import "fmt"

func changeValue(x *string) {
	*x = "test"
}
func changeValue2(x string) {
	x = "oto test"
	fmt.Println("this isn't a new value:", x)
}
func main() {
	a := "pruu"
	fmt.Println("func main:", a)
	changeValue(&a)
	fmt.Println("func main:", a)
	fmt.Println("------------------------------------------------------")
	a = "two"
	fmt.Println("func main:", a)
	changeValue2(a)
	fmt.Println("func main:", a)

}
