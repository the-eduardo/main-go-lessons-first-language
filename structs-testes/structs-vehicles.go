package main

import "fmt"

/*- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.*/
type vehicle struct {
	portas uint8
	cor    string
}
type caminhonete struct {
	vehicle
	traçãoNasQuatro bool
}
type sedan struct {
	vehicle
	modeloLuxo bool
}

func main() {
	c1 := caminhonete{
		vehicle: vehicle{
			portas: 6,
			cor:    "azul",
		},
		traçãoNasQuatro: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			portas: 4,
			cor:    "azul",
		}}
	s2 := sedan{
		vehicle: vehicle{
			portas: 4,
			cor:    "azul",
		},
		modeloLuxo: true,
	}
	fmt.Println(c1)
	fmt.Println(s1)
	fmt.Println(s2)
}
