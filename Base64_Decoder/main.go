package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

//create a function that encode base 64 string to string
func encodeBase64(s string) string {
	//encode the string to base64
	b64 := base64.StdEncoding.EncodeToString([]byte(s))
	//return the encoded string

	return b64
}
func decodeBase64(s string) string {
	//decode the string from base64
	b64, err := base64.StdEncoding.DecodeString(s)
	//check for error
	if err != nil {
		fmt.Println(err)
	}
	//return the decoded string
	return string(b64)
}

// user able to choose to encode or decode
type choice struct {
	encode bool
	decode bool
}

func (choice) getChoice() choice {
	var c choice
	fmt.Println("\t 1. Encode")
	fmt.Println("\t 2. Decode")
	fmt.Println("\t 3. Exit")
	fmt.Println("Enter your choice:")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println(err)
	}
	switch choice {
	case 1:
		c.encode = true
		c.decode = false
	case 2:
		c.encode = false
		c.decode = true
	case 3:
		os.Exit(0)
	default:
		os.Exit(0)
	}
	return c
}
func main() {
	//get the choice from user
	fmt.Println("Do you want to encode or decode?")
	c := choice{}.getChoice()
	//create a scanner
	scanner := bufio.NewScanner(os.Stdin)
	//get the input from user
	fmt.Println("Enter the string:")
	scanner.Scan()
	//get the input string
	s := scanner.Text()
	//check for the choice
	if c.encode {
		//encode the string
		fmt.Println(encodeBase64(s))
	} else if c.decode {
		//decode the string
		fmt.Println(decodeBase64(s))
	} else {
		//exit the program
		os.Exit(0)
	}
	rAgain()
}
func rAgain() {
	main()
}
