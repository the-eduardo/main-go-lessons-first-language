package main

import "fmt"

type person struct {
	name string
	age  int
}
type secretAgent struct {
	person
	working bool
	salary  int
}
type baker struct {
	person
	cakesmade int
	bakery    string
}
type human interface {
	Greetings()
}

func oi(h human) {
	switch h.(type) {
	case person:
		h.Greetings()
	case secretAgent:
		fmt.Println("")
		h.Greetings()
		fmt.Printf("\t I'm a secret agent, and I'm working %t, and I earn %d\n", h.(secretAgent).working, h.(secretAgent).salary)
	case baker:
		fmt.Println("")
		h.Greetings()
		fmt.Println("\t I'm a baker, and I make", h.(baker).cakesmade, "cakes a day, and I own", h.(baker).bakery)
	}
}

func (p person) Greetings() {
	fmt.Println("Hi, I'm", p.name, "with", p.age, "years old.")
}
func (s secretAgent) Greetings() {
	fmt.Println("Hi, I'm", s.name, "with", s.age, "years old")
}
func (b baker) Greetings() {
	fmt.Println("Hi, I'm", b.name, "with", b.age, "years old")
}

func main() {
	p := person{name: "John", age: 40}
	s := secretAgent{person: person{name: "Anna", age: 35}, working: true, salary: 15000}
	b := baker{person: person{name: "Mone", age: 20}, cakesmade: 20, bakery: "Baker's Bakery"}
	oi(p)
	oi(s)
	oi(b)
	fmt.Println("------------------------------------------------------")
	b.Greetings()
	p.Greetings()
	s.Greetings()

}
