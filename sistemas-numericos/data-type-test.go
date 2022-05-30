package main
import ("fmt"
  "runtime")

func main() {
  s := "hello world!"
  sb := []byte(s)
  for _, v := range sb {
    fmt.Printf("%v | %b | %T | %#U | %#x\n", v, v, v, v, v)
  }
  a := 78745
  fmt.Printf("\n\n\n %d \t %b \t %#x", a ,a ,a )
  fmt.Println("\n", runtime.GOOS, "|", runtime.GOARCH)
}
/* https://pkg.go.dev/fmt
Integer:

%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
**/
