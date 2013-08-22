package input

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

const IntMax = int(^uint(0) >> 1)
const IntMin = -IntMax - 1

// Line reading closure, to save on bufio.NewReader overhead.
// Prompts with a supplied string.
func GetReader() func(string) (string) {
    in := bufio.NewReader(os.Stdin)
    return func(p string) (string) {
        fmt.Print(p)
        s, _ := in.ReadString('\n')
        s = strings.TrimRight(s, "\n\r")
        return s
    }
}

//Type parser func shorthand...
type Parser func(string, interface{}) error

//Parses a line from stdin with specified func,
//retrying indefinitely on failure
func Read(p string,
          f func(string, interface{}) error,
          x interface{}) {
    in := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(p)
        s, e := in.ReadString('\n') //REPLACE _
        s = strings.TrimRight(s, "\n\r")
        if e = f(s, x); e == nil {
            break
        }
        fmt.Println(e)
    }
}
//Function to parse ints, for use with input.Read()
func Int(s string, x interface{}) error {
    i, e := strconv.Atoi(s)
    if e != nil {
        return fmt.Errorf("Not an integer")
    }
    *x.(*int) = i
    return nil
}

// Prompts with p until a valid int is entered
func ReadInt(p string) int {
    return ReadRangedInt(p, IntMin, IntMax)
}

func ReadRangedInt(p string, min, max int) int {
    in := GetReader()
    for {
        s := in(p)
        i, e := strconv.Atoi(s)
        if e == nil {
            if i >= min && i <= max {
                return i
            }
            fmt.Printf("Enter a number between %d & %d\n", min, max)
        } else {
            fmt.Println("Not an integer")
        }
    }
}
