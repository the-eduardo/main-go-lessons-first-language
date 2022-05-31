package main
import "fmt"
func main() {
	for h := 0; h < 3; h++ {
		for m := 0; m < 60; m++{
			for s := 0; s < 60; {
				s++
				{
					fmt.Printf("%vh%vm%vs\n", h, m, s)
				}}}}}
