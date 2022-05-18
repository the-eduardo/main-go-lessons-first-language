package main
import ("fmt")

  func square(x, y float64) (float64, float64){
    return (x*x), (y*4)
  }

  func main () {
    var x float64
    var abc string

    fmt.Printf("Please enter content:")
    fmt.Scan(&x)
    y := x
    area, perimetro := square(x, y)
    fmt.Print("Your has square area is " area, " meters\n", "And has a perimeter of ", perimetro, " meters")

  }
