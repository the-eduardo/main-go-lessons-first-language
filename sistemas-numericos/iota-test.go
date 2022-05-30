package main
import "fmt"

type Stereotype int

const (
    TypicalNoob Stereotype = iota // 0
    TypicalHipster                // 1
    TypicalUnixWizard             // 2
    TypicalStartupFounder         // 3
)

func SoSayethThe(character Stereotype) string {
    var s string
    switch character {
    case TypicalNoob:
        s = "I'm a confused ninja rockstar."
    case TypicalHipster:
        s = "Everything was better we programmed uphill and barefoot in the snow on the SUTX 5918"
    case TypicalUnixWizard:
        s = "sudo grep awk sed %!#?!!1!"
    case TypicalStartupFounder:
        s = "exploit compelling convergence to syndicate geo-targeted solutions"
    }
    return s
}


type ByteSize float64

const (
    _           = iota                   // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (iota * 10) // 1 << (1 * 10)
    MB                                   // 1 << (10*2)
    GB                                   // 1 << (10*3)
    TB                                   // 1 << (10*4)
    PB                                   // 1 << (10*5)
    EB                                   // 1 << (10*6)
    ZB                                   // 1 << (10*7)
    YB                                   // 1 << (10*8)
)
func bsize() {
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%v\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%v\n", MB)
  fmt.Printf("%b\t", GB)
	fmt.Printf("%v\n", GB)
  fmt.Printf("%b\t", TB)
  fmt.Printf("%v\n", TB)
  fmt.Printf("%b\t", PB)
  fmt.Printf("%v\n", PB)
  fmt.Printf("%b\t", EB)
  fmt.Printf("%v\n", EB)
  fmt.Printf("%b\t", ZB)
  fmt.Printf("%v\n", ZB)
  fmt.Printf("%b\t", YB)
  fmt.Printf("%v\n", YB)
}



func main() {

    fmt.Println(SoSayethThe(1))
    fmt.Println("----------------------------------------------")
    bsize()

}
