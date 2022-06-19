package main

import "fmt"

type person struct {
	name  string
	idade int
}
type professional struct {
	person
	title  string
	income int
}

func main() {
	person1 := person{
		"robson",
		27,
	}
	person2 := person{
		"adalberto",
		46,
	}
	pro1 := professional{
		person{"julio", 28}, "programador", 2800,
	}
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(pro1)
	aaa()
}

func aaa() {
	x := struct {
		name string
		pass int
	}{name: "Mairo",
		pass: 555559}
	fmt.Println(x)
}
