package main
import "fmt"
import "time"

  func main() {
  var x string
  x = "teste"
  fmt.Println(x, time.Now())

  var yy string
  fmt.Scan(&yy)
  if x != yy {
    println("x != y")
  } else {
    println("x is equal to y")
  }
  }
