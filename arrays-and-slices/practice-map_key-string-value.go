package main

import "fmt"

func main() {
	pessoa := map[string][]string{
		"joao_aaa":       {"Dormir", "comer", "comer", "Dormir"},
		"amelio_brandão": {"cantar", "desenhar"},
		"julio_souza":    {"dançar", "codar"},
		"paulo_ferras":   {"codar", "desenhar"},
		"thiago_otashi":  {"desenhar", "voar"},
	}
	pessoa["this_is"] = []string{"a_test", "Ler", "Cozinhar"}
	delete(pessoa, "joao_aaa")
	for a, b := range pessoa {
		fmt.Println(a)
		for i, h := range b {
			fmt.Println("\t", i, "|", h)
		}
	}

}
