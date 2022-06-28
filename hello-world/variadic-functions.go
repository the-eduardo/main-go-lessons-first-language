package main

import "fmt"

func main() {
	slic := []int{1, 2, 3, 4, 500}
	a, b := variadicFunction(slic...)
	fmt.Println("Total =", a, "Length =", b)
	a, b = variadicFunction()
	fmt.Println("Total =", a, "Length =", b)
}
func variadicFunction(a ...int) (int, int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum, len(a)
}

/* Output:
Total = 510 Length = 5
Total = 0 Length = 0
*/
