package main

import "fmt"

func main() {
	x := struct {
		nome    string
		sabor   string
		ondetem []string
		bemcom  map[string]string
	}{
		nome:    "geléia",
		sabor:   "melancia",
		ondetem: []string{"casa", "trabalho", "escola"},
		bemcom: map[string]string{
			"casa":     "breakfast",
			"trabalho": "café",
			"escola":   "comida",
		},
	}
	//fmt.Println(x)
	fmt.Println(x.nome, x.sabor, ":")
	fmt.Println("\t tem em:")
	for _, a := range x.ondetem {
		fmt.Println("\t\t-", a)
	}
	fmt.Println("\t vai bem com:")
	for _, a := range x.bemcom {
		fmt.Println("\t\t-", a)
	}
}
