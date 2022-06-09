package main

import "fmt"

func main() {
	ssmultid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("\n", ssmultid)
	fmt.Println("\n", ssmultid[1][0])
	fmt.Println("-----------------------------------------------")
	feia()
}

func feia() {
	ss := [][][][]int{

		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				[]int{10, 20, 30, 40, 50},
			},
		},

		[][][]int{
			[][]int{
				[]int{2, 4, 6, 8, 10},
			},
			[][]int{
				[]int{3},
			},
		},
	}
	fmt.Println(ss)
	fmt.Println(ss[1][0][0][2])
}
