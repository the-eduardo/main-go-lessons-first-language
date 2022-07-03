package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() {
	result := q.lado * q.lado
	fmt.Println("Área do quad. é:", result)
}
func (c circulo) area() {
	result := math.Pi * c.raio * c.raio
	fmt.Println("Área do circ. é:", result)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()

}
func main() {
	a := quadrado{lado: 10}
	b := circulo{raio: 6.12}

	a.area()
	b.area()

	fmt.Println("------------------------------------------------------")
	medida(a)
	medida(b)
}
