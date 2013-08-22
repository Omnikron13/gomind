package score

import "fmt"
import "os"
import "bufio"

import "github.com/Omnikron13/gomind/code"

type Score struct {
    c code.Code //Code played
    t int       //Tries
    p string    //Player
}

//Create a new Score object from a Code, try count & player name
func New(c code.Code, t int, p string) (s Score) {
    s = Score{c, t, p}
    return
}

//Formats a Score object into a meaningful score for output
func (s Score) String() (str string) {
    scr, d := s.c.Score(s.t)
    str = fmt.Sprintf("%.0f\t%.0f\t%v", d, scr, s.p)
    return
}

func (s Score) Append() {
    f, e := os.OpenFile(scoreFile, os.O_RDWR|os.O_APPEND, 0660)
    defer f.Close()
    if e != nil {
        fmt.Println(e)
    }
    buf := bufio.NewWriter(f)
    buf.WriteString(s.marshal())
    buf.WriteRune('\n')
    buf.Flush()
}

//Generate a savable representation of a Score object
func (s Score) marshal() (str string) {
    str = fmt.Sprintf("%v %d %d %v",
                      s.c,
                      s.c.Depth(),
                      s.t,
                      s.p)
    return
}

//Process a saved score string into a Score object
func unmarshal(str string) Score {
    //Code, player
    var cs, p string
    //Depth, 'tries'
    var d, t int
    fmt.Sscanf(str, "%s %d %d %s", &cs, &d, &t, &p)
    c := code.Load(cs, d)
    return New(c, t, p)
}
