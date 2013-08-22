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
          f Parser,
          x interface{}) {
    in := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(p)
        s, e := in.ReadString('\n')
        s = strings.TrimRight(s, "\n\r")
        if e = f(s, x); e == nil {
            break
        }
        fmt.Println(e)
    }
}

//Returns a Parser for ints within specified range
func IntParser(min, max int) Parser {
    return func(s string, x interface{}) error {
        i, e := strconv.Atoi(s)
        if e != nil {
            return fmt.Errorf("Must be an integer")
        }
        if i < min || i > max {
            return fmt.Errorf("Must be in range %d-%d", min, max)
        }
        *x.(*int) = i
        return nil
    }
}

//Potentially useful timesaver for when a return
//is preferably to passing by refernce
func ReadInt(p string, min, max int) (i int) {
    Read(p, IntParser(min, max), &i)
    return
}
