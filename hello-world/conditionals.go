package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Add your code below:
	rand.Seed(time.Now().UnixNano())

	amountLeft := rand.Intn(10000)

	fmt.Println("amountLeft is: ", amountLeft)

	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
	fmt.Println(time.Now().UnixNano())

	if lessonLearned := true; lessonLearned {
		fmt.Println("Great job! You can continue on to the next exercise.")
	} else {
		fmt.Println("Practice makes perfect.")
	}
	// Create more conditions below!
	color := "f"
	switch color {
	case "yellow":
		fmt.Println("amarelo")
	case "blue":
		fmt.Println("azul")
	case "red":
		fmt.Println("rojo")
	default:
		fmt.Println("nothing")

		number := rand.Intn(100)
		fmt.Println("ramdom number:", number)
	}

}
