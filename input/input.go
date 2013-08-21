package input

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

// Line reading closure, to save on bufio.NewReader overhead.
// Prompts with a supplied string.
func getReader() func(string) (string) {
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
    in := getReader()
    for {
        s := in(p)
        i, e := strconv.Atoi(s)
        if e == nil {
            return i
        }
        fmt.Println("Not an integer")
    }
}
