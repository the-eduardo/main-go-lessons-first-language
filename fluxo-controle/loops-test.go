package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 20; i++ {
		if i == 3 {
			// faz isso quando o número não é par
			continue
		}
		fmt.Printf("%v ", i)
	}

	fmt.Println("\n----------------------------------------------------")
	secnd()
	fmt.Printf("|")
	fmt.Printf("\n----------------------------------------------------\n")
	old()

}

func secnd() {

	x := 0

	for x <= 21 {

		x++

		if x%2 != 0 {
			continue
		}

		fmt.Printf("| %v ", x)

	}

}

/*- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.*/
func old() {
	ano := 1980
	x := 1540
	for {
		if x == ano {
			break
		}
		fmt.Printf("%v // ", x)
		x = x + 10

	}
}
