package score

import "fmt"
import "bytes"
import "os"
import "bufio"
import "sort"

//Number of scores displayed
const numShown  = 10  -1 //-1 for ease of comparison...
//Amount internal slice is grown by
const capInc    = 10
//File scores are written to & read from
const scoreFile = "scores"

type Board []Score

//Load saved scores from the file 'scores'
func Load() Board {
    f, _ := os.Open(scoreFile)
    buf := bufio.NewReader(f)
    b := make(Board, 0, capInc)
    for i := 0; ; i++ {
        if ln, e := buf.ReadString('\n'); e == nil {
            s := unmarshal(ln)
            if len(b) == cap(b) {
                tmp := make(Board, len(b), cap(b)+capInc)
                copy(tmp, b)
                b = tmp
            }
            b = b[:i+1]
            b[i] = s
        } else {
            break
        }
    }
    sort.Sort(b)
    return b
}

//Format Board object into a scoreboard suitable for output
func (b Board) String() string {
    var buf bytes.Buffer
    buf.WriteString("Diff.\tScore\tPlayer\n")
    buf.WriteString("-----\t-----\t------\n")
    for i, s := range b {
        //Limit scores to Top <numShown>...
        if i > numShown {
            break
        }
        buf.WriteString(fmt.Sprintln(s))
    }
    return buf.String()
}

//sort.Sort() compliance..
func (b Board) Len() int {
    return len(b)
}
func (b Board) Less(x, y int) bool {
    i, _ := b[x].c.Score(b[x].t)
    j, _ := b[y].c.Score(b[y].t)
    return i > j
}
func (b Board) Swap(x, y int) {
    b[x], b[y] = b[y], b[x]
}
