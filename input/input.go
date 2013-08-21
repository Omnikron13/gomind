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
