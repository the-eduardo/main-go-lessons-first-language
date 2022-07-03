package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}
type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	fmt.Printf("Rectangle area of %v: ", r)
	return r.width * r.height
}
func (c *circle) area() float64 {
	fmt.Printf("Circle area of %v: ", c)
	return math.Pi * c.radius * c.radius
}
func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{5}
	r1 := rect{5, 7}
	c2 := circle{10}
	r2 := rect{10, 20}
	for _, shape := range []shape{&c1, &c2, &r1, &r2} {
		fmt.Println(getArea(shape))
	}
}
