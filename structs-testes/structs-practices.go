package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sorvete   []string
}

func main() {
	p1 := pessoa{"joao", "alberto", []string{"melancia", "laranja", "jaca"}}
	p2 := pessoa{"rodrigo", "manhatan", []string{"melancia", "jabuticaba"}}
	p3 := pessoa{"p3", "p3", []string{"abacate", "pera", "maçã"}}
	//	fmt.Println(p1)
	//	fmt.Println(p2)
	//	fmt.Println(p3)
	fmt.Println(p1.nome, p1.sobrenome, ":")
	for _, a := range p1.sorvete {
		fmt.Println("\t", a)
	}
	fmt.Println(p2.nome, p2.sobrenome, ":")
	for _, a := range p2.sorvete {
		fmt.Println("\t", a)
	}
	fmt.Println(p3.nome, p3.sobrenome, ":")
	for _, a := range p3.sorvete {
		fmt.Println("\t", a)
	}

}
