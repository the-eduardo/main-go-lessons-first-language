package main

import "fmt"

func GPA(midtermGrade, finalGrade float32) (string, float32) {

	averageGrade := (midtermGrade + finalGrade) / 2

	var gradeLetter string

	if averageGrade > 90 {
		gradeLetter = "A, aproved"
	} else if averageGrade > 80 {
		gradeLetter = "B, aproved"
	} else if averageGrade > 70 {
		gradeLetter = "C, aproved"
	} else if averageGrade > 60 {
		gradeLetter = "D, desaproved"
	} else {
		gradeLetter = "F, desaproved"
	}

	return gradeLetter, averageGrade
}

func main() {
	var mid, final float32
	mid = 81.4
	final = 74.9
	var myAverage float32
	var myGrade string
	myGrade, myAverage = GPA(mid, final)
	fmt.Println("score: ", myAverage, myGrade)
}
