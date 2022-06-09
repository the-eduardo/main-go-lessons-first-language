package main

import "fmt"

func main() {
	fmt.Println("-----------------------------------------------")
	friends := map[string]int{
		"abc": 1,
		"def": 2,
		"ghi": 3,
	}
	fmt.Println(friends)
	fmt.Println("-----------------------------------------------")
	total := 0
	for key, value := range friends {
		fmt.Println(key, value)
		total += value
	}
	fmt.Println(total)

	fmt.Println("-----------------------------------------------")
	fmt.Println(friends["abc"])
	friends["jkl"] = 4
	fmt.Println(friends)
	fmt.Println("-----------------------------------------------")
	// comma ok idiom
	testando, ok := friends["teste"]
	fmt.Println(testando, ok)
	fmt.Println("-----------------------------------------------")
	delete(friends, "def")
	fmt.Println(friends) // deleted element from map
	fmt.Println("-----------------------------------------------")
}
