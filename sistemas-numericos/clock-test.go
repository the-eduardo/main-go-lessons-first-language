package main
import ("fmt")

func main() {
  h :=0
  m :=0
  s :=0
  for h < 3 {
  for m < 60 {
    s ++
    fmt.Printf("%vh%vm%vs\n", h, m, s)
    if s > 59 {
      m ++
      s = 0
      fmt.Printf("%vh%vm%vs\n", h, m, s)
    }
  }
  if m > 59 {
    m=0
    h++
  }
      }
}
