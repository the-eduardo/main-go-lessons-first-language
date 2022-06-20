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
	p3 := pessoa{"nome", "sobrenome", []string{"abacate", "pera", "maçã"}}
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
	fmt.Println("---------------------------------------------")
	mape()

}
func mape() {
	pessoas := make(map[string]pessoa)
	pessoas["joão"] = pessoa{"joão", "alberto", []string{"melancia", "laranja", "jaca"}}
	pessoas["ervilha"] = pessoa{"tomás", "segundo", []string{"abacate", "pera", "maçã"}}
	//fmt.Println(pessoas)
	for _, valor := range pessoas {
		fmt.Println("")
		fmt.Println("Meu nome é", valor.nome, valor.sobrenome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sorvete {
			fmt.Println("–", valor)
		}
	}
}
