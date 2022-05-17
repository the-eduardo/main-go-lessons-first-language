// https://www.youtube.com/watch?v=094y1Z2wpJg&t=1098s

//a number is odd *3+1
// if even /2

package main
import ("fmt"
    "math/rand"
    "time"
  )

  func oddEven(num int) int{
   if num & 1 != 0 {
      fmt.Println("ODD")
      num = num*3+1
      fmt.Println(num)
      return num
   } else {
      fmt.Println("EVEN")
      num = num/2
      fmt.Println(num)
      return num
   }
}

func main(){
  var sum int
  rand.Seed(time.Now().UnixNano())

  oddEven(rand.Intn(100))
  oddEven(rand.Intn(100))
  oddEven(rand.Intn(100))

  sum = (rand.Intn(100))
  oddEven(sum)
  if sum & 1 != 0 {
     fmt.Println("ODD ***")
     sum = sum*3+1
     fmt.Println(sum)
  } else {
     fmt.Println("EVEN ---")
     sum = sum/2
     fmt.Println(sum)
  }

}
