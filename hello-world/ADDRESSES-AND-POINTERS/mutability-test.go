package main

import "fmt"

type Artists struct {
	Name, Genre string
	Songs       int
}

func newReleases(a *Artists) int { //passing an Artist by reference
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artists{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newReleases(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
