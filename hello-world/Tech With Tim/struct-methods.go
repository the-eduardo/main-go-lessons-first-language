package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Scores []int
}

func (s Student) avgGrande() float64 {
	var sum int
	for _, score := range s.Scores {
		sum += score
	}
	return float64(sum) / float64(len(s.Scores))
}
func (s *Student) setAge(age int) {
	s.Age = age
}
func (s Student) maxScore() int {
	sum := 0
	for _, v := range s.Scores {
		if v > sum {
			sum = v
		}
	}
	return sum
}

func main() {
	s := Student{
		Name:   "John",
		Age:    20,
		Scores: []int{79, 80, 43, 84, 85},
	}
	fmt.Printf("%v is %v years old \nAnd has an average score of %v and max score of %v", s.Name, s.Age, s.avgGrande(), s.maxScore())

}
