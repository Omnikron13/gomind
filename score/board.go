package score

import "fmt"
import "bytes"
import "os"
import "bufio"

type Board []Score

//Load saved scores from the file 'scores'
func Load() Board {
    f, _ := os.Open("scores")
    buf := bufio.NewReader(f)
    b := make(Board, 0, 10)
    for i := 0; ; i++ {
        if ln, e := buf.ReadString('\n'); e == nil {
            s := unmarshal(ln)
            if len(b) == cap(b) {
                tmp := make(Board, len(b), cap(b)+10)
                copy(tmp, b)
                b = tmp
            }
            b = b[:i+1]
            b[i] = s
        } else {
            break
        }
    }
    b.bsort()
    return b
}

//Format Board object into a scoreboard suitable for output
func (b Board) String() string {
    var buf bytes.Buffer
    buf.WriteString("Diff.\tScore\tPlayer\n")
    buf.WriteString("-----\t-----\t------\n")
    for i, s := range b {
        //Limit scores to Top 10...
        if i > 9 {
            break
        }
        buf.WriteString(fmt.Sprintln(s))
    }
    return buf.String()
}

//Rough implementation of a bubble sort...
func (b Board) bsort() {
    for done := false; !done; done = true {
        for x := len(b)-1; x > 0; x-- {
            s1, _ := b[x].c.Score(b[x].t)
            s2, _ := b[x-1].c.Score(b[x].t)
            if s1 > s2 {
                b[x], b[x-1] = b[x-1], b[x]
                done = false
            }
        }
    }
}
