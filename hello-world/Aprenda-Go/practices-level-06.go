package main

import "fmt"

func main() {
	fon := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fon2 := []int{12, 3, 48, 9, 110}
	fmt.Println(funcion(fon...))
	fmt.Println(otrafuncion(fon))
	fmt.Println(otrafuncion(fon2))
}
func funcion(x ...int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}
func otrafuncion(x []int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}

/*- Exercício:
  - Crie uma função que retorne um int
  - Crie outra função que retorne um int e uma string
  - Chame as duas funções
  - Demonstre seus resultados

- Exercício:
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.



*/
