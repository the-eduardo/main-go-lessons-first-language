package main

import "fmt"

var z interface{}

func zz() {
	switch z.(type) {
	case bool:
		fmt.Println("bool |", z)
		z = 22
		zz()
	case int:
		fmt.Println("int |", z)
		z = 2.2
		zz()
	case float64:
		fmt.Println("float64 |", z)
		z = "22"
		zz()
	case string:
		fmt.Println("string |", z)
	default:
		fmt.Println("fon |", z)
		z = true
		zz()
	}
}

func main() {
	x := 0
	switch {
	case x > 10, x == 1:
		fmt.Println("x > 10 or x = 1")
	case x == 10:
		fmt.Println("x é 10")
		fallthrough
	case (x < 10):
		fmt.Println("x < 10 ou é 10")
	default:
		fmt.Println("fon")

	}
	fmt.Println("-------------------------")
	zz()
}
