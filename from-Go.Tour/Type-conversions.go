/**------------------------------------------------------
Type conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.
-----------------------------------------------------**/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) //f aqui é float64
	z := uint(f)                                  // Var f é float64 porém var Z é uint
	fmt.Println(x, y, z)
	fmt.Printf("Tipos %T f %T x, %T y, %T z Valores: %v %v %v %v \n", f, x, y, z, f, x, y, z)
	// print: Tipos float64 f int x, int y, uint z Valores: 5 3 4 5
	//                                       ^^^^^

	fmt.Println("--------------------------------")
	inference()
}

func inference() {
	i := 420          // int
	j := i            // j is an int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("j is %v of type %T\n", j, j)
	fmt.Printf("i is %v of type %T\n", i, i)
	fmt.Printf("f is %v of type %T\n", f, f)
	fmt.Printf("g is %v of type %T\n", g, g)
}
      /*---------------Prints
      j is 420 of type int
      i is 420 of type int
      f is 3.142 of type float64
      g is (0.867+0.5i) of type complex128
      ----------------------------------------*/
