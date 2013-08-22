// gomind is a version of the 'Mastermind' board game
// written in Go.
package main

import "fmt"
import "os"

import "github.com/Omnikron13/gomind/input"
import "github.com/Omnikron13/gomind/code"
import "github.com/Omnikron13/gomind/score"

const Version = "gomind-[alpha]"

func main() {
    io := input.GetReader()
    for {
        switch s := io("play, scores, quit? "); s {
        case "p", "play":
            game(setup())
        case "s", "scores":
            b := score.Load()
            fmt.Print(b)
        case "q", "quit", "exit":
            os.Exit(0)
        case "h", "help", "?":
            help()
        case "v", "version":
            fmt.Println(Version)
        case "debug":
            c := setup()
            fmt.Println(c)
            game(c)
        default:
            fmt.Println("Command not understood... (Try h[elp])")
        }
    }
}

func setup() (c code.Code) {
    l := input.ReadInt("Code Length [2-8]: ", 2, 8)
    d := input.ReadInt("Code Depth [2-16]: ", 2, 16)
    c = code.New(l, d)
    return
}

func game(c code.Code) {
    //fmt.Println(c.String()) //Cheap debug...
    fmt.Printf("Length %d, Depth %d...\n", c.Length(), c.Depth())
    for i := 1;; i++ {
        p := fmt.Sprintf("%d: ", i)
        t := c.Read(p)
        b, w := c.Try(t)
        fmt.Printf("Black; %d, White; %d\n", b, w)
        //Check for the win...
        if b == c.Length() {
            fmt.Println("You win!")
            s, d := c.Score(i)
            fmt.Printf("Difficulty: %.0f\n", d)
            fmt.Printf("Score: %.0f\n", s)
            save(c, i)
            break
        }
    }
}

//Save score mechanic...
func save(c code.Code, t int) {
    io := input.GetReader()
    for {
        switch s := io("Save this score [y/n]? "); s {
        case "yes", "y":
            p := io("Enter your name: ")
            scr := score.New(c, t, p)
            scr.Append()
            fmt.Println("Score saved!")
            return
        case "no", "n":
            return
        default:
            fmt.Println("You must answer y[es] or [n]")
        }
    }
}

//Prints command list & explanation
func help() {
    //Basics
    fmt.Println("[Basic Commands]")
    fmt.Print(" h[elp]    - ")
    fmt.Print("display this list\n")
    fmt.Print(" p[lay]    - ")
    fmt.Print("start a new game with custom length/depth\n")
    fmt.Print(" s[cores]  - ")
    fmt.Print("list the (top 10) scores\n")
    fmt.Print(" q[uit]    - ")
    fmt.Print("exit gomind\n")
    //Extra
    fmt.Println("[More]")
    fmt.Print(" v[ersion] - ")
    fmt.Print("display version details\n")
}
