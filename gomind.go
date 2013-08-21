// gomind is a version of the 'Mastermind' board game
// written in Go.
package main 

import "fmt"

import "github.com/Omnikron13/gomind/input"

func main() {
    for {
        i := input.readInt("Int: ")
        fmt.Printf("%d*3=%d\n", i, i*3)
    }
}
