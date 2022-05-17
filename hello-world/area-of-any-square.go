package main
import ("fmt")

  func square(x, y float64) (float64, float64){
    return (x*x), (y*4)
  }

  func main () {
    var x float64

    fmt.Printf("Please enter content:")
    fmt.Scan(&x)
    y := x
    area, perimetro := square(x, y)
    fmt.Println("Your has square area is", area,"and has a perimeter of", perimetro)

  }
