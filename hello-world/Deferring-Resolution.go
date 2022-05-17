/*Deferring Resolution
We can delay a function call to the end of the current scope by
using the defer keyword. defer tells Go to run a function,
but at the end of the current function.

This is useful for logging, file writing, and other utilities.
*/

package main
import "fmt"

func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!") //!important
  taxRate := .06143
  fmt.Println("Calculating Taxes...")

  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  }


  taxValue := (revenue - (deductions * credits)) * taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}

func main() {
  var rev, ded, cre float64
  rev = 100000
  ded = 100
  cre = 100
  var final float64
   final = calculateTaxes(rev, ded, cre)
  fmt.Println(final)

}

/* In the above example, when we call calculateTaxes() we
immediately defer a message, "Taxes Calculated!".
This does not print until the end of the function
(after the taxes have been calculated and are about to be returned).
Normally, we would consider adding fmt.Println("Taxes Calculated!")
at the end of calculateTaxes(). But, we have multiple return
statements in our code, instead of adding a print statement right
before each return, we use defer and it prints regardless of when
our function ends.*/

/*
Calculating Taxes...
Taxes Calculated!
      5528.7
[Finished in 0.453s]*/




//--------------------------------------------------------------------------------
 //package main
//import "fmt"

func queryDatabase(query string) string {
  var result string
  connectDatabase()
  // Add deferred call to disconnectDatabase() here
  defer disconnectDatabase()
  if query == "SELECT * FROM coolTable;" {
    result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
  }
  fmt.Println(result)
  return result
}

func connectDatabase() {
  fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
  fmt.Println("Disconnecting from the database.")
}

func main2() {
  queryDatabase("SELECT * FROM coolTable;")
}
