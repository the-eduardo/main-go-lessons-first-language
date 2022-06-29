package main

import "fmt"

type persona struct {
	name, nick string
	age        int
}

func (p persona) sayhi() {
	fmt.Println("Hi, I'm", p.name, p.nick, ",I'm", p.age, "years old.")
}
func main() {

	a := persona{name: "John", nick: "Paulman", age: 40}

	a.sayhi()
}
