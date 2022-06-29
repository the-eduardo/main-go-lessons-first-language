package main

import "fmt"

type person struct {
	name, nick string
	age        int
}

func (p person) sayhi() {
	fmt.Println("Hi, I'm", p.name, p.nick, ",I'm", p.age, "years old.")
}
func main() {

	a := person{name: "John", nick: "Paulman", age: 40}

	a.sayhi()
}
