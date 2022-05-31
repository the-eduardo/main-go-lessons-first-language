// -------------------------- OLD ----------------------------------

/* package main
import (
	"math"
  "fmt"
)

// Define specialComputation() here

func main() {
  var a, b, c, d float64
  a = .0214
  b = 1.02
  c = 0.312
  d = 4.001

  // Replace the following four lines with specialComputation()
  a = math.Log2(math.Sqrt(math.Tanh(a)))
  b = math.Log( math.Sqrt(math.Tanh(b)))
  c = math.Log(math.Sqrt( math.Tan(c) ))
  d = math.Log2(math.Sqrt(math.Tanh(d)))

  fmt.Println(a, b, c, d)
}*/

// -------------------------- MAIN ----------------------------------
package main

import (
	"fmt"
	"math"
)

// Define specialComputation() here
func specialComputation(x float64) float64 {
	return math.Log2(math.Sqrt(math.Tan(x)))
}

func main() {
	var a, b, c, d float64
	a = .0214
	b = 1.02
	c = 0.312
	d = 4.001

	// Replace the following four lines with specialComputation()
	a = specialComputation(a)
	b = specialComputation(b)
	c = specialComputation(c)
	d = specialComputation(d)

	fmt.Println(a, b, c, d)
}
