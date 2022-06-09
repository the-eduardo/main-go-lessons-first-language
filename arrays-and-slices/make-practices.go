package main

import "fmt"

/*func main() {
	var (
		estados = map[string]int{
			"Acre":                1,
			"Alagoas":             2,
			"Amapá":               3,
			"Amazonas":            4,
			"Bahia":               5,
			"Ceará":               6,
			"Espírito Santo":      7,
			"Goiás":               8,
			"Maranhão":            9,
			"Mato Grosso":         10,
			"Mato Grosso do Sul":  11,
			"Minas Gerais":        12,
			"Pará":                13,
			"Paraíba":             14,
			"Paraná":              15,
			"Pernambuco":          16,
			"Piauí":               17,
			"Rio de Janeiro":      18,
			"Rio Grande do Norte": 19,
			"Rio Grande do Sul":   20,
			"Rondônia":            21,
			"Roraima":             22,
			"Santa Catarina":      23,
			"São Paulo":           24,
			"Sergipe":             25,
			"Tocantins":           26,
		}
	)
	fmt.Println(estados)
	for estado, codigo := range estados {
		fmt.Println(estado, codigo)
	}
}*/
func main() {
	estados := make([]string, 26, 26)
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	fmt.Println("len, cap =", len(estados), cap(estados))
	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
	fmt.Println("-----------------------------------------------")
	address()
}
func address() {
	slices := [][]string{
		[]string{"Acre", "Alagoas", "Amapá"},
		[]string{"Amazonas", "Bahia", "Ceará"},
		[]string{"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais"},
		[]string{"Paraíba", "Pará", "Pernambuco"},
	}
	for _, valor := range slices {
		fmt.Println(valor[0])
		for _, item := range valor[1:] {
			fmt.Println("\t", item)
		}
	}
}
