package main
import "fmt"
func main() {
  aaa := []string{"total", "1st", "2nd", "3rd", "4th", "5th", "6th", "7th","8th"}
  win := aaa[1:4]
  los := aaa[4:len(aaa)]
  fmt.Println(win)
  fmt.Println(los)
  fmt.Println(aaa[:], "\n----------------")
  for i :=0; i<len(aaa); i++ {
    fmt.Println(aaa[i])
  }
  aaa = append(aaa[1:3], aaa[6:]...)
  fmt.Println(aaa)
}
