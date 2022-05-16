package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

eludedGuards := rand.Intn(100)
//First Conditional (Sneak past guards)
isHeistOn := true

if eludedGuards >= 50 && isHeistOn {
  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  fmt.Println("isHeistOn is currently:", isHeistOn)
  fmt.Println("********************************")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
    fmt.Println("isHeistOn is currently:", isHeistOn)
    fmt.Println("********************************")
  }
  var openedVault = rand.Intn(100)
  if isHeistOn && openedVault >= 70 {
    fmt.Println("Is open! Grab and GO!")
    fmt.Println("********************************")
} else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can’t be opened")
  fmt.Println("isHeistOn is currently:", isHeistOn)
  fmt.Println("********************************")
  }
  if isHeistOn && openedVault >= 70 {
  var leftSafely = rand.Intn(5)
  switch leftSafely {
        case 0:
        isHeistOn = false;
        fmt.Println("Looks like you tripped an alarm... run?   #")
        case 1:
        isHeistOn = false
        fmt.Println("Turns out vault doors don't open from the inside...")
        case 2:
        isHeistOn = false
        fmt.Println("When did they start raising dogs in vaults??")
        case 3:
        fmt.Println("We are out! Move!")
        default:
        fmt.Println("Start the getaway car!")

     }} else {
       isHeistOn = false;
     }

     if isHeistOn {
        var amtStolen = 10000 + rand.Intn(1000000)
        fmt.Println("We have $", amtStolen)
        fmt.Println("********************************")
    }
    fmt.Println(" Heist is currently: ", isHeistOn)
}
