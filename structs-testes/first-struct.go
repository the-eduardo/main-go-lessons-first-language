package main

import "fmt"

type school struct {
	name     string
	nickname string
	grade    uint8
	approved bool
}

func main() {
	student1 := school{
		name:     "rodolfo",
		nickname: "testado",
		grade:    8,
		approved: true,
	}
	student2 := school{
		name:     "joana",
		nickname: "silva",
		grade:    2,
		approved: false,
	}
	student3 := school{
		name:     "roberto",
		nickname: "abcd",
		grade:    10,
		approved: true,
	}
	// fmt.Println student 1, 2, 3
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)
}
